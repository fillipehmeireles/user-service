package schemas

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID          int `gorm:"primaryKey"`
	Name        string
	Email       string    `gorm:"unique"`
	PhoneNumber string    `gorm:"unique"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
	DeletedAt   gorm.DeletedAt
}
