package usecase

import (
	"clip/repo"
	"fmt"
)

type clipboard struct {
	repo *repo.Clipboard
}

type Clipboard interface {
	SaveInClipboard(data []byte) error
}

func NewClipboard(repo *repo.Clipboard) Clipboard {
	return &clipboard{}
}

func (c *clipboard) SaveInClipboard(data []byte) error {
	err, lastData := c.repo.(data)
	fmt.Printf("last: %v\n",lastData)
	return err
}
