package models

import (
	"time"

	"gorm.io/gorm"
)

// User Model with created_at and updated_at fields
type User struct {
	gorm.Model
	ID        uint           `json:"id" gorm:"primaryKey"` // gorm:"primaryKey" is used to set the primary key
	Username  string         `json:"username" gorm:"unique"`
	Password  string         `json:"password"`
	Name      string         `json:"name"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-"`

	Appointments []Appointment `json:"appointments"`
}

// TableName sets the table name
func (User) TableName() string {
	return "users"
}

// UserResponse is the response for the user model after login
type UserResponse struct {
	Username       string `json:"username"`
	Name           string `json:"name"`
	Token          string `json:"token"`
	ExpirationTime int64  `json:"expiration_time"`
}

// UserLogin is the request for the user login
type UserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// UserRegister is the request for the user register
type UserRegister struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
}
