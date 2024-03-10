package models

import (
	"time"

	"gorm.io/gorm"
)

// Patient Model with created_at and updated_at fields
type Patient struct {
	gorm.Model
	ID                  uint           `json:"id" gorm:"primaryKey"` // gorm:"primaryKey" is used to set the primary key
	MedicalRecordNumber string         `json:"medical_record_number" gorm:"unique"`
	Name                string         `json:"name"`
	Gender              uint           `json:"gender"`
	BirthDate           string         `json:"birth_date"`
	Address             string         `json:"address"`
	PhoneNumber         string         `json:"phone_number"`
	CreatedAt           time.Time      `json:"-"`
	UpdatedAt           time.Time      `json:"-"`
	DeletedAt           gorm.DeletedAt `json:"-"`

	Appointments []Appointment `json:"appointments"`
}

// PatientResponse is the response for the patient model
type PatientResponse struct {
	ID                  uint   `json:"id"`
	MedicalRecordNumber string `json:"medical_record_number"`
	Name                string `json:"name"`
	Gender              uint   `json:"gender"`
	BirthDate           string `json:"birth_date"`
	Address             string `json:"address"`
	PhoneNumber         string `json:"phone_number"`
}
