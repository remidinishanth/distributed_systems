---
layout: page
title: "Lab 1: MapReduce Implementation (MIT 6.5840)"
category: "google_map_reduce"
---

## Lab 1: MapReduce

Lab URL: [https://pdos.csail.mit.edu/6.824/labs/lab-mr.html](https://pdos.csail.mit.edu/6.824/labs/lab-mr.html)

Build a MapReduce system: a **coordinator** that hands out tasks to workers and handles failures, and **workers** that process map and reduce tasks, communicating via RPC.

### Architecture Overview

```
                    ┌──────────────┐
                    │  Coordinator │
                    │  (1 process) │
                    └──────┬───────┘
                           │ RPC (Unix socket)
              ┌────────────┼────────────┐
              │            │            │
        ┌─────┴──────┐ ┌───┴──────┐ ┌───┴──────┐
        │  Worker 1  │ │ Worker 2 │ │ Worker 3 │
        └────────────┘ └──────────┘ └──────────┘
```

### Files Modified

Only three files: `mr/rpc.go`, `mr/coordinator.go`, `mr/worker.go`

Entry points (`main/mrcoordinator.go`, `main/mrworker.go`) are unchanged. The worker loads Map/Reduce functions at runtime via Go's plugin system (`-buildmode=plugin`).

### Two-Phase Execution

```
Phase 1: MAP                          Phase 2: REDUCE
─────────────────                     ─────────────────
Input files → Map workers             Intermediate files → Reduce workers
  pg-*.txt      │                       mr-X-Y files           │
                ▼                                              ▼
         Intermediate files                          Output files
         mr-X-Y (JSON encoded)                       mr-out-Y
         X = map task ID                              Y = reduce task ID
         Y = reduce bucket (ihash(key) % nReduce)
```

- All map tasks must complete before any reduce task begins.
- Each map task produces `nReduce` intermediate files.
- Each reduce task reads from all map outputs for its partition.

### Key Design Decisions

#### 1. RPC Protocol

Two RPCs:
- **GetTask**: Worker asks coordinator for work. Reply contains task type (Map/Reduce/Wait/Exit), task ID, filename, nReduce, nMap.
- **ReportTask**: Worker tells coordinator a task is done.

```go
const (
    MapTask    = 0  // Execute map function
    ReduceTask = 1  // Execute reduce function
    WaitTask   = 2  // No task available, sleep and retry
    ExitTask   = 3  // All work done, worker should exit
)
```

#### 2. Coordinator State Machine

Each task goes through states: **Idle → InProgress → Completed**

```go
type TaskInfo struct {
    State     int       // Idle / InProgress / Completed
    StartTime time.Time // When task was assigned (for timeout detection)
}
```

The coordinator tracks:
- `mapTasks[]` and `reduceTasks[]` — per-task state
- `allMapDone` — gate between map and reduce phases
- `allDone` — signals `Done()` to shut down coordinator

#### 3. Fault Tolerance (10-second timeout)

When assigning tasks, the coordinator checks for timed-out workers:

```go
if t.State == TaskIdle ||
   (t.State == TaskInProgress && time.Since(t.StartTime) > 10*time.Second) {
    // Assign this task
}
```

No need for heartbeats or worker IDs — just reassign tasks that take too long. This handles both slow and crashed workers.

#### 4. Atomic File Writes

Workers write to temp files, then `os.Rename()` atomically:

```go
tmpFile, _ := ioutil.TempFile("", "mr-tmp-*")
// ... write all data ...
tmpFile.Close()
os.Rename(tmpFile.Name(), finalName)
```

This prevents partial output from being visible to other workers or the reduce phase.

#### 5. Intermediate File Format

JSON-encoded KeyValue pairs, one per line:

```go
// Write (map)
enc := json.NewEncoder(file)
enc.Encode(&kv)

// Read (reduce)
dec := json.NewDecoder(file)
dec.Decode(&kv)
```

File naming: `mr-{mapTaskID}-{reduceTaskID}`

#### 6. Key Partitioning

```go
bucket := ihash(kv.Key) % nReduce
```

Uses FNV hash to deterministically assign keys to reduce partitions. All occurrences of the same key go to the same reducer.

#### 7. Worker Lifecycle

```
Worker starts
    │
    ▼
┌─► GetTask RPC ──► MapTask? ──► Read file, mapf(), write mr-X-Y, ReportTask
│       │
│       ├──► ReduceTask? ──► Read mr-*-Y, sort, reducef(), write mr-out-Y, ReportTask
│       │
│       ├──► WaitTask? ──► sleep 100ms
│       │
│       └──► ExitTask? ──► return
│       │
│       └──► RPC failed? ──► return (coordinator gone)
│
└───────┘ (loop)
```

Workers exit gracefully when the coordinator shuts down (RPC dial fails → return instead of log.Fatal).

### Concurrency

- Coordinator uses a single `sync.Mutex` protecting all state.
- Workers are independent processes — no shared memory.
- Multiple workers can execute map (or reduce) tasks in parallel.

### Test Results

All 7 tests pass with `-race`:

| Test | What it checks |
|------|---------------|
| `TestWc` | Word count correctness (3 workers) |
| `TestIndexer` | Inverted index correctness (2 workers) |
| `TestMapParallel` | Map tasks run concurrently |
| `TestReduceParallel` | Reduce tasks run concurrently |
| `TestJobCount` | Each map task runs exactly once |
| `TestEarlyExit` | No premature exit before output is finalized |
| `TestCrashWorker` | Correct output despite random worker crashes |
