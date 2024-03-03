package models

import (
	"time"

	"gorm.io/gorm"
)

// Appointment Data Type Model with created_at and updated_at fields
type AppointmentDataType struct {
	gorm.Model
	ID        uint      `json:"id" gorm:"primaryKey"` // gorm:"primaryKey" is used to set the primary key
	Name      string    `json:"name"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

// Appointment Data Type Status Enum
const (
	OBSTETRICS = "OBSTETRICS" // for Kebidanan
	// SOAP - for Kandungan
	SUBJECTIVE = "SUBJECTIVE"
	OBJECTIVE  = "OBJECTIVE"
	ASSESSMENT = "ASSESSMENT"
	PLAN       = "PLAN"
)
