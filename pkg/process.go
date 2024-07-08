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
	DeleteClipboardLastDayData() error
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
			fmt.Printf("text: %v\n", string(text))
			p.clipboard.SaveInClipboard(text)
		}
		fmt.Printf("image: %v\n", reader.ReadImage())
		time.Sleep(1 * time.Second)
	}
}

func (p *process) DeleteClipboardLastDayData() error {
	now:=time.Now()

	// TODO: change time to 00:00:00
	lastDay:=now.AddDate(0,0,-1).UTC()
	fmt.Printf("Deleting data from %v\n",lastDay)
	err := p.clipboard.DeleteClipboardData(lastDay)
	if err != nil {
		fmt.Println("Error while deleting the data")
		return err
	}
	return nil
}