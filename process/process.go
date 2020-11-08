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
	Id      int
	Counter int
	Iterate bool
	Chanel  *bool
}

func (process *Process) Start() {
	for {
		if process.Iterate == false {
			break
		}
		process.Counter++
		if *process.Chanel == true {
			fmt.Printf("%d : %d \n", process.Id, process.Counter)
			time.Sleep(time.Second * 1)
		}
	}
}
func (process *Process) Kill() {
	process.Iterate = false
}
