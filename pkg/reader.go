package pkg

import (
	"golang.design/x/clipboard"
)

type reader struct {}

type Reader interface {
	ReadText() []byte
	ReadImage() []byte
}


func NewReader() Reader {
	err:=clipboard.Init()
	if err != nil {
		panic(err)
	}
	return &reader{}
}

func (r *reader) ReadText() []byte {
	c:=clipboard.Read(clipboard.FmtText)
	return c
}

func (r *reader) ReadImage() []byte {
	i:=clipboard.Read(clipboard.FmtImage,)
	return i
}