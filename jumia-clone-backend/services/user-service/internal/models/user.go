package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// User represents a user in the system
type User struct {
	ID        string         `gorm:"type:uuid;primary_key" json:"id"`
	FirstName string         `gorm:"type:varchar(100);not null" json:"first_name"`
	LastName  string         `gorm:"type:varchar(100);not null" json:"last_name"`
	Email     string         `gorm:"type:varchar(255);uniqueIndex;not null" json:"email"`
	Password  string         `gorm:"type:varchar(255);not null" json:"-"` // Never return password in JSON
	Phone     string         `gorm:"type:varchar(20)" json:"phone"`
	Address   string         `gorm:"type:text" json:"address"`
	IsActive  bool           `gorm:"default:true" json:"is_active"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// BeforeCreate hook to generate UUID before creating user
func (u *User) BeforeCreate(tx *gorm.DB) error {
	if u.ID == "" {
		u.ID = uuid.New().String()
	}
	return nil
}

// TableName specifies the table name for the User model
func (User) TableName() string {
	return "users"
}
