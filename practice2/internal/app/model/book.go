package model

import (
	"gorm.io/gorm"
	"time"
)

type Book struct {
	ID          uint64         `json:"-" gorm:"AUTO_INCREMENT;PRIMARY_KEY;not null"`
	UID         string         `json:"uid" gorm:"uniqueIndex;type:varchar(36);not null"`
	Title       *string        `json:"title" gorm:"not null"`
	Author      *string        `json:"author"`
	Description *string        `json:"description"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}
