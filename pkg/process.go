package pkg

import (
	"fmt"
	"time"
)


type process struct {}
type Process interface {
	Init()
}

func NewProcess() Process {
	return &process{}
}

func (p *process ) Init() {
	reader:=NewReader()
	for{
		fmt.Printf("text: %v\n",string(reader.ReadText()))
		fmt.Printf("image: %v\n",reader.ReadImage())
		time.Sleep(1*time.Second)
	}
}