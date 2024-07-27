package usecase

import (
	"clip/database"
	"clip/logger"
	"clip/pkg/config"
	"clip/repo"
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type clipboard struct {
	repo repo.Clipboard
	l    logger.Logger
	C    *config.Config
}

type Clipboard interface {
	SaveInClipboard(data []byte, dataType database.Datatype) error
	DeleteClipboardData(date time.Time) error
	GetLast10() (error, []database.Clipboard)
	SearchInClipboard(data string) (error, []database.Clipboard)
}

func NewClipboard(repo repo.Clipboard, l logger.Logger, c *config.Config) Clipboard {
	return &clipboard{
		repo,
		l,
		c,
	}
}

func (c *clipboard) SaveInClipboard(data []byte, dataType database.Datatype) error {
	err, lastStoredData := c.repo.LastStoredData()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = c.repo.Insert(data, database.Clipboard{}, dataType)
			c.l.Error(fmt.Sprintf("Error while getting last stored data %v", err))
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

func (c *clipboard) SearchInClipboard(data string) (error, []database.Clipboard) {
	err, searchResult := c.repo.Search(data)
	if err != nil {
		c.l.Error(fmt.Sprintf("Error while searching data %v", err))
		return err, searchResult
	}
	return nil, searchResult
}
