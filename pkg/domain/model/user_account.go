package model

import (
	"time"
)

// UserAccount Struct
type UserAccount struct {
	ID        int `gorm:"primaryKey"`
	CreatedAt time.Time
	CreatedBy int
	UpdatedAt time.Time
	UpdatedBy int
	Name      string
	Email     string
	Password  string
}
