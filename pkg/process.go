package pkg

import (
	"clip/database"
	"clip/logger"
	"clip/usecase"
	"fmt"
	"time"
)

type process struct {
	clipboard usecase.Clipboard
	logger    logger.Logger
}
type Process interface {
	Init()
	DeleteClipboardLastDayData() error
}

func NewProcess(u usecase.Clipboard, l logger.Logger) Process {
	return &process{
		u,
		l,
	}
}

func (p *process) Init() {
	reader := NewReader()
	for {
		text := reader.ReadText()
		if len(text) != 0 || text != nil {
			p.clipboard.SaveInClipboard(text, database.Text)
		}
		i := reader.ReadImage()
		if i != nil || len(i) != 0 {
			p.clipboard.SaveInClipboard(i, database.Image)
		}
		time.Sleep(1 * time.Second)
	}
}

func (p *process) DeleteClipboardLastDayData() error {
	now := time.Now()

	// TODO: change time to 00:00:00
	lastDay := now.AddDate(0, 0, -1).UTC()
	p.logger.Info(fmt.Sprintf("Deleting data before %v", lastDay))
	err := p.clipboard.DeleteClipboardData(lastDay)
	if err != nil {
		p.logger.Error(fmt.Sprintf("Error while deleting data before %v", err))
		return err
	}
	return nil
}
