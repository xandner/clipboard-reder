package repo

import (
	"clip/database"
	"time"

	"gorm.io/gorm"
)

type clipboard struct {
	db *gorm.DB
}

type Clipboard interface {
	Insert(ctx []byte, e database.Clipboard) error
	FindByContent(ctx []byte) (error, database.Clipboard)
	DeleteFromDate(date time.Time) error
}

func NewClipboard(db *gorm.DB) Clipboard {
	return &clipboard{
		db,
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

func (c *clipboard) Insert(ctx []byte, e database.Clipboard) error {
	err := c.db.Create(&database.Clipboard{Data: ctx}).Error
	return err
}

func (c *clipboard) DeleteFromDate(date time.Time) error {
	err := c.db.Where("created_at < ?", date).Unscoped().Delete(&database.Clipboard{}).Error
	return err
}