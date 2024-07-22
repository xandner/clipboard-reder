package database

import (
	"time"

	"gorm.io/gorm"
)

type Datatype string

const (
	Text  Datatype = "text"
	Image Datatype = "image"
)

type Clipboard struct {
	gorm.Model
	Datatype  Datatype `gorm:"not null" json:"type"`
	Data      []byte   `gorm:"not null" json:"data"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
