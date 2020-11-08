package process

import (
	"fmt"
	"time"
)

type System struct {
	QueueProcess []QueueProcess
}
type QueueProcess interface {
	Start()
	Kill()
}
type Process struct {
	Id         int
	Counter    int
	Iterate    bool
	NewProcess bool
}

func (process *Process) Start() {
	for {
		if process.Iterate == true {
			process.Counter++
			fmt.Printf("%d : %d \n", process.Id, process.Counter)
			time.Sleep(time.Millisecond * 500)
		}
	}
}
func (process *Process) Kill() {
	process.Iterate = false
}
