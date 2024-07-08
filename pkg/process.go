package pkg

import (
	"clip/usecase"
	"fmt"
	"time"
)

type process struct {
	clipboard usecase.Clipboard
}
type Process interface {
	Init()
}

func NewProcess(u usecase.Clipboard) Process {
	return &process{
		u,
	}
}

func (p *process) Init() {
	reader := NewReader()
	for {
		text := reader.ReadText()
		if len(text) != 0 || text != nil {
			p.clipboard.SaveInClipboard(text)
		}
		fmt.Printf("image: %v\n", reader.ReadImage())
		time.Sleep(1 * time.Second)
	}
}
