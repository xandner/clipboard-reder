package usecase

import (
	"clip/database"
	"clip/repo"
	"fmt"
	"time"
)

type clipboard struct {
	repo repo.Clipboard
}

type Clipboard interface {
	SaveInClipboard(data []byte,dataType database.Datatype) error
	DeleteClipboardData(date time.Time) error
}

func NewClipboard(repo repo.Clipboard) Clipboard {
	return &clipboard{
		repo,
	}
}

func (c *clipboard) SaveInClipboard(data []byte, dataType database.Datatype) error {
	err, lastStoredData := c.repo.LastStoredData()
	if err != nil {
		if err.Error() == "record not found" {
			err = c.repo.Insert(data, database.Clipboard{},dataType)
			return err
		}
		fmt.Println("Error while finding the data")
		return err
	}
	if string(lastStoredData.Data) != string(data) {
		err = c.repo.Insert(data, database.Clipboard{}, dataType)
		return err
	}
	return nil
}

func (c *clipboard) DeleteClipboardData(date time.Time) error {
	err := c.repo.DeleteFromDate(date)
	if err != nil {
		fmt.Println("Error while deleting the data")
		return err
	}
	return nil
}
