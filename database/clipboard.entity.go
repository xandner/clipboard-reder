package database

import (
	"time"

	"gorm.io/gorm"
)

type Datatype string

const (
	Text Datatype = "text"
	Image Datatype = "image"
)

type Clipboard struct {
	gorm.Model
	Datatype  Datatype   `gorm:"not null"`
	Data      []byte     `gorm:"not null"`
	CreatedAt time.Time
    UpdatedAt time.Time
}