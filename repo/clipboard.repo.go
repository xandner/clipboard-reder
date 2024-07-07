package repo

import (
	"clip/database"
	"context"

	"gorm.io/gorm"
)

type clipboard struct {
	db *gorm.DB
}

type Clipboard interface {
	Insert(ctx context.Context, e database.Clipboard) error
	FindByContent(ctx []byte) (error, database.Clipboard)
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

func (c *clipboard) Insert(ctx context.Context, e database.Clipboard) error {
	err := c.db.WithContext(ctx).Create(&e).Error
	return err
}
