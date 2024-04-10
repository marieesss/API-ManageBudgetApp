package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        uint
	Email     string `gorm:"uniqueIndex; column:email"`
	Name      string `gorm:"type:varchar(255);not null"`
	Surname   string `gorm:"type:varchar(255);not null"`
	Password  string `gorm:"type:varchar(255);not null"`
	Birthday  *time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}
