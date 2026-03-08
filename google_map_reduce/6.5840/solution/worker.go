package mr

import "fmt"
import "log"
import "net/rpc"
import "hash/fnv"
import "os"
import "io/ioutil"
import "encoding/json"
import "sort"
import "time"

// Map functions return a slice of KeyValue.
type KeyValue struct {
	Key   string
	Value string
}

// for sorting by key.
type ByKey []KeyValue

func (a ByKey) Len() int           { return len(a) }
func (a ByKey) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByKey) Less(i, j int) bool { return a[i].Key < a[j].Key }

// use ihash(key) % NReduce to choose the reduce
// task number for each KeyValue emitted by Map.
func ihash(key string) int {
	h := fnv.New32a()
	h.Write([]byte(key))
	return int(h.Sum32() & 0x7fffffff)
}

var coordSockName string // socket for coordinator

// main/mrworker.go calls this function.
func Worker(sockname string, mapf func(string, string) []KeyValue,
	reducef func(string, []string) string) {

	coordSockName = sockname

	for {
		args := GetTaskArgs{}
		reply := GetTaskReply{}
		ok := call("Coordinator.GetTask", &args, &reply)
		if !ok {
			// Coordinator is gone, exit
			return
		}

		switch reply.TaskType {
		case MapTask:
			doMap(mapf, reply.TaskID, reply.File, reply.NReduce)
		case ReduceTask:
			doReduce(reducef, reply.TaskID, reply.NMap)
		case WaitTask:
			time.Sleep(100 * time.Millisecond)
		case ExitTask:
			return
		}
	}
}

func doMap(mapf func(string, string) []KeyValue, taskID int, filename string, nReduce int) {
	// Read input file
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("cannot open %v", filename)
	}
	content, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalf("cannot read %v", filename)
	}
	file.Close()

	// Call map function
	kva := mapf(filename, string(content))

	// Partition into nReduce buckets
	buckets := make([][]KeyValue, nReduce)
	for _, kv := range kva {
		bucket := ihash(kv.Key) % nReduce
		buckets[bucket] = append(buckets[bucket], kv)
	}

	// Write intermediate files using temp files for atomicity
	for i := 0; i < nReduce; i++ {
		oname := fmt.Sprintf("mr-%d-%d", taskID, i)
		tmpFile, err := ioutil.TempFile("", "mr-tmp-*")
		if err != nil {
			log.Fatalf("cannot create temp file: %v", err)
		}
		enc := json.NewEncoder(tmpFile)
		for _, kv := range buckets[i] {
			if err := enc.Encode(&kv); err != nil {
				log.Fatalf("cannot encode kv: %v", err)
			}
		}
		tmpFile.Close()
		os.Rename(tmpFile.Name(), oname)
	}

	// Report completion
	reportArgs := ReportTaskArgs{TaskType: MapTask, TaskID: taskID}
	reportReply := ReportTaskReply{}
	call("Coordinator.ReportTask", &reportArgs, &reportReply)
}

func doReduce(reducef func(string, []string) string, taskID int, nMap int) {
	// Read all intermediate files for this reduce task
	var intermediate []KeyValue
	for i := 0; i < nMap; i++ {
		iname := fmt.Sprintf("mr-%d-%d", i, taskID)
		file, err := os.Open(iname)
		if err != nil {
			continue // file may not exist if map produced no output for this partition
		}
		dec := json.NewDecoder(file)
		for {
			var kv KeyValue
			if err := dec.Decode(&kv); err != nil {
				break
			}
			intermediate = append(intermediate, kv)
		}
		file.Close()
	}

	// Sort by key
	sort.Sort(ByKey(intermediate))

	// Write output using temp file for atomicity
	oname := fmt.Sprintf("mr-out-%d", taskID)
	tmpFile, err := ioutil.TempFile("", "mr-out-tmp-*")
	if err != nil {
		log.Fatalf("cannot create temp file: %v", err)
	}

	// Call reduce for each distinct key
	i := 0
	for i < len(intermediate) {
		j := i + 1
		for j < len(intermediate) && intermediate[j].Key == intermediate[i].Key {
			j++
		}
		values := []string{}
		for k := i; k < j; k++ {
			values = append(values, intermediate[k].Value)
		}
		output := reducef(intermediate[i].Key, values)
		fmt.Fprintf(tmpFile, "%v %v\n", intermediate[i].Key, output)
		i = j
	}

	tmpFile.Close()
	os.Rename(tmpFile.Name(), oname)

	// Report completion
	reportArgs := ReportTaskArgs{TaskType: ReduceTask, TaskID: taskID}
	reportReply := ReportTaskReply{}
	call("Coordinator.ReportTask", &reportArgs, &reportReply)
}

// example function to show how to make an RPC call to the coordinator.
//
// the RPC argument and reply types are defined in rpc.go.
func CallExample() {

	// declare an argument structure.
	args := ExampleArgs{}

	// fill in the argument(s).
	args.X = 99

	// declare a reply structure.
	reply := ExampleReply{}

	// send the RPC request, wait for the reply.
	// the "Coordinator.Example" tells the
	// receiving server that we'd like to call
	// the Example() method of struct Coordinator.
	ok := call("Coordinator.Example", &args, &reply)
	if ok {
		// reply.Y should be 100.
		fmt.Printf("reply.Y %v\n", reply.Y)
	} else {
		fmt.Printf("call failed!\n")
	}
}

// send an RPC request to the coordinator, wait for the response.
// usually returns true.
// returns false if something goes wrong.
func call(rpcname string, args interface{}, reply interface{}) bool {
	// c, err := rpc.DialHTTP("tcp", "127.0.0.1"+":1234")
	c, err := rpc.DialHTTP("unix", coordSockName)
	if err != nil {
		return false
	}
	defer c.Close()

	if err := c.Call(rpcname, args, reply); err == nil {
		return true
	}
	return false
}
