package database

import (
	"time"

	"gorm.io/gorm"
)

type datatype string

const (
	Text datatype = "text"
	Image datatype = "image"
)

type Clipboard struct {
	gorm.Model
	Datatype  datatype   `gorm:"not null"`
	Data      []byte     `gorm:"not null"`
	CreatedAt time.Time
    UpdatedAt time.Time
}
