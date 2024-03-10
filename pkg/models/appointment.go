package models

import (
	"time"

	"gorm.io/gorm"
)

// Appointment Model with created_at and updated_at fields
type Appointment struct {
	gorm.Model
	ID                uint           `json:"id" gorm:"primaryKey"` // gorm:"primaryKey" is used to set the primary key
	PatientID         uint           `json:"patient_id"`
	UserID            uint           `json:"user_id"` // User is doctor who created the appointment
	AppointmentDate   string         `json:"appointment_date"`
	AppointmentTypeID uint           `json:"appointment_type_id"`
	CreatedAt         time.Time      `json:"-"`
	UpdatedAt         time.Time      `json:"-"`
	DeletedAt         gorm.DeletedAt `json:"-"`

	AppointmentType AppointmentType   `json:"appointment_type"`
	AppointmentData []AppointmentData `json:"appointment_data"`
}

// AppointmentResponse is the response for the appointment model. Appointment Type Name is included
type AppointmentResponse struct {
	ID                uint   `json:"id"`
	PatientID         uint   `json:"patient_id"`
	UserID            uint   `json:"user_id"`
	AppointmentDate   string `json:"appointment_date"`
	AppointmentTypeID uint   `json:"appointment_type_id"`
	AppointmentType   string `json:"appointment_type"`
}

// AppointmentCreate is the request for the appointment create
type AppointmentCreate struct {
	PatientID         uint   `json:"patient_id"`
	UserID            uint   `json:"user_id"`
	AppointmentDate   string `json:"appointment_date"`
	AppointmentTypeID uint   `json:"appointment_type_id"`
}
