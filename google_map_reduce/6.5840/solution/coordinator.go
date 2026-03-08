package mr

import "log"
import "net"
import "os"
import "net/rpc"
import "net/http"
import "sync"
import "time"

const (
	TaskIdle       = 0
	TaskInProgress = 1
	TaskCompleted  = 2
)

type TaskInfo struct {
	State     int
	StartTime time.Time
}

type Coordinator struct {
	mu          sync.Mutex
	files       []string
	nReduce     int
	nMap        int
	mapTasks    []TaskInfo
	reduceTasks []TaskInfo
	allMapDone  bool
	allDone     bool
}

// GetTask is the RPC handler for workers requesting tasks.
func (c *Coordinator) GetTask(args *GetTaskArgs, reply *GetTaskReply) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	reply.NReduce = c.nReduce
	reply.NMap = c.nMap

	if !c.allMapDone {
		// Try to assign a map task
		for i, t := range c.mapTasks {
			if t.State == TaskIdle || (t.State == TaskInProgress && time.Since(t.StartTime) > 10*time.Second) {
				reply.TaskType = MapTask
				reply.TaskID = i
				reply.File = c.files[i]
				c.mapTasks[i].State = TaskInProgress
				c.mapTasks[i].StartTime = time.Now()
				return nil
			}
		}
		// All map tasks assigned but not all done yet
		reply.TaskType = WaitTask
		return nil
	}

	// Map phase done, try reduce tasks
	for i, t := range c.reduceTasks {
		if t.State == TaskIdle || (t.State == TaskInProgress && time.Since(t.StartTime) > 10*time.Second) {
			reply.TaskType = ReduceTask
			reply.TaskID = i
			c.reduceTasks[i].State = TaskInProgress
			c.reduceTasks[i].StartTime = time.Now()
			return nil
		}
	}

	if c.allDone {
		reply.TaskType = ExitTask
	} else {
		reply.TaskType = WaitTask
	}
	return nil
}

// ReportTask is the RPC handler for workers reporting task completion.
func (c *Coordinator) ReportTask(args *ReportTaskArgs, reply *ReportTaskReply) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if args.TaskType == MapTask {
		c.mapTasks[args.TaskID].State = TaskCompleted
		// Check if all map tasks are done
		allDone := true
		for _, t := range c.mapTasks {
			if t.State != TaskCompleted {
				allDone = false
				break
			}
		}
		c.allMapDone = allDone
	} else if args.TaskType == ReduceTask {
		c.reduceTasks[args.TaskID].State = TaskCompleted
		// Check if all reduce tasks are done
		allDone := true
		for _, t := range c.reduceTasks {
			if t.State != TaskCompleted {
				allDone = false
				break
			}
		}
		c.allDone = allDone
	}
	return nil
}

// an example RPC handler.
//
// the RPC argument and reply types are defined in rpc.go.
func (c *Coordinator) Example(args *ExampleArgs, reply *ExampleReply) error {
	reply.Y = args.X + 1
	return nil
}

// start a thread that listens for RPCs from worker.go
func (c *Coordinator) server(sockname string) {
	rpc.Register(c)
	rpc.HandleHTTP()
	os.Remove(sockname)
	l, e := net.Listen("unix", sockname)
	if e != nil {
		log.Fatalf("listen error %s: %v", sockname, e)
	}
	go http.Serve(l, nil)
}

// main/mrcoordinator.go calls Done() periodically to find out
// if the entire job has finished.
func (c *Coordinator) Done() bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.allDone
}

// create a Coordinator.
// main/mrcoordinator.go calls this function.
// nReduce is the number of reduce tasks to use.
func MakeCoordinator(sockname string, files []string, nReduce int) *Coordinator {
	c := Coordinator{
		files:       files,
		nReduce:     nReduce,
		nMap:        len(files),
		mapTasks:    make([]TaskInfo, len(files)),
		reduceTasks: make([]TaskInfo, nReduce),
	}

	c.server(sockname)
	return &c
}
