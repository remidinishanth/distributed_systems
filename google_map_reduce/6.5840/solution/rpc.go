package mr

//
// RPC definitions.
//
// remember to capitalize all names.
//

//
// example to show how to declare the arguments
// and reply for an RPC.
//

type ExampleArgs struct {
	X int
}

type ExampleReply struct {
	Y int
}

// Task types
const (
	MapTask    = 0
	ReduceTask = 1
	WaitTask   = 2
	ExitTask   = 3
)

// GetTaskArgs is sent by a worker to request a task.
type GetTaskArgs struct {
}

// GetTaskReply is the coordinator's response with a task assignment.
type GetTaskReply struct {
	TaskType int
	TaskID   int
	File     string // input file for map tasks
	NReduce  int
	NMap     int
}

// ReportTaskArgs is sent by a worker to report task completion.
type ReportTaskArgs struct {
	TaskType int
	TaskID   int
}

// ReportTaskReply is the coordinator's acknowledgment.
type ReportTaskReply struct {
}
