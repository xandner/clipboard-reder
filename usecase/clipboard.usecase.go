package usecase

import (
	"clip/database"
	"clip/logger"
	"clip/repo"
	"fmt"
	"time"
)

type clipboard struct {
	repo repo.Clipboard
	l    logger.Logger
}

type Clipboard interface {
	SaveInClipboard(data []byte, dataType database.Datatype) error
	DeleteClipboardData(date time.Time) error
	GetLast10() (error, []database.Clipboard)
}

func NewClipboard(repo repo.Clipboard, l logger.Logger) Clipboard {
	return &clipboard{
		repo,
		l,
	}
}

func (c *clipboard) SaveInClipboard(data []byte, dataType database.Datatype) error {
	err, lastStoredData := c.repo.LastStoredData()
	if err != nil {
		if err.Error() == "record not found" {
			err = c.repo.Insert(data, database.Clipboard{}, dataType)
			return err
		}
		c.l.Error(fmt.Sprintf("Error while getting last stored data %v", err))
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
		c.l.Error(fmt.Sprintf("Error while deleting the data %v", err))
		return err
	}
	return nil
}

func (c *clipboard) GetLast10() (error, []database.Clipboard) {
	err, data := c.repo.GetLast10()
	if err != nil {
		c.l.Error(fmt.Sprintf("Error while getting last 10 data %v", err))
		return err, data
	}
	return nil, data
}