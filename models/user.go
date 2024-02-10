package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

// gorm.Model definition
type GormModel struct {
	ID        uuid.UUID      `gorm:"type:uuid;primaryKey"`
	CreatedAt time.Time      `gorm:"index" json:"created" `
	UpdatedAt time.Time      `json:"updated"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted"`
}

type User struct {
	GormModel
	Email    string `gorm:"uniqueIndex"`
	Password string
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	// Set the CreatedAt and UpdatedAt timestamps
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
	return nil
}

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	// Update the UpdatedAt timestamp
	u.UpdatedAt = time.Now()
	return nil
}
