package repo

import (
	"clip/database"
	"clip/logger"
	"time"

	"gorm.io/gorm"
)

type clipboard struct {
	db *gorm.DB
	l  logger.Logger
}

type Clipboard interface {
	Insert(ctx []byte, e database.Clipboard, dataType database.Datatype) error
	FindByContent(ctx []byte) (error, database.Clipboard)
	DeleteFromDate(date time.Time) error
	LastStoredData() (error, database.Clipboard)
}

func NewClipboard(db *gorm.DB,l logger.Logger) Clipboard {
	return &clipboard{
		db,
		l,
	}
}

func (c *clipboard) FindByContent(ctx []byte) (error, database.Clipboard) {
	var clipboardData database.Clipboard
	err := c.db.Where("data = ?", ctx).First(&clipboardData).Error
	if err != nil {
		return err, clipboardData
	}
	return nil, clipboardData
}

func (c *clipboard) Insert(ctx []byte, e database.Clipboard, dataType database.Datatype) error {
	err := c.db.Create(&database.Clipboard{Data: ctx, Datatype: dataType}).Error
	return err
}

func (c *clipboard) DeleteFromDate(date time.Time) error {
	err := c.db.Where("created_at < ?", date).Unscoped().Delete(&database.Clipboard{}).Error
	return err
}

func (c *clipboard) LastStoredData() (error, database.Clipboard) {
	var clipboardData database.Clipboard
	err := c.db.Last(&clipboardData).Error
	if err != nil {
		return err, clipboardData
	}
	return nil, clipboardData
}
