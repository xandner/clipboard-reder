package usecase

import (
	"clip/database"
	"clip/repo"
	"fmt"
)

type clipboard struct {
	repo repo.Clipboard
}

type Clipboard interface {
	SaveInClipboard(data []byte) error
}

func NewClipboard(repo repo.Clipboard) Clipboard {
	return &clipboard{
		repo,
	}
}

func (c *clipboard) SaveInClipboard(data []byte) error {
	err, _ := c.repo.FindByContent(data)
	if err != nil {
		if err.Error() == "record not found" {
			err = c.repo.Insert(data, database.Clipboard{})
			return nil
		}
		fmt.Println("Error while finding the data")
		return err
	}
	return nil
}
