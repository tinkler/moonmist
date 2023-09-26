package gm

import (
	"time"

	"gorm.io/gorm"
)

type Model struct {
	CreatedAt time.Time `gorm:"->"`
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
