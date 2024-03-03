package models

import (
	"time"

	"gorm.io/gorm"
)

// Appointment Type Model with created_at and updated_at fields
// Status is used to set the status of the appointment type, value get from enum (ACTIVE, INACTIVE, DELETED)
type AppointmentType struct {
	gorm.Model
	ID        uint      `json:"id" gorm:"primaryKey"` // gorm:"primaryKey" is used to set the primary key
	Name      string    `json:"name"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`

	Appointments []Appointment `json:"appointments"`
}

// Appointment Type Status Enum
const (
	ACTIVE   = "ACTIVE"
	INACTIVE = "INACTIVE"
	DELETED  = "DELETED"
)

// AppointmentTypeResponse is the response for the appointment type model
type AppointmentTypeResponse struct {
	ID     uint   `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

// AppointmentTypeCreate is the request for the appointment type create. Status is set to ACTIVE by default
type AppointmentTypeCreate struct {
	Name string `json:"name"`
}
