package models

import (
	"time"

	"gorm.io/gorm"
)

// Appointment Data Type Model with created_at and updated_at fields
type AppointmentDataType struct {
	gorm.Model
	ID        uint           `json:"id" gorm:"primaryKey"` // gorm:"primaryKey" is used to set the primary key
	Name      string         `json:"name"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeleteAt  gorm.DeletedAt `json:"-"`
}

// Appointment Data Type Name Enum
const (
	ADT_Obstetrics = "OBSTETRICS" // for Kebidanan
	// SOAP - for Kandungan
	ADT_Subjective = "SUBJECTIVE"
	ADT_Objective  = "OBJECTIVE"
	ADT_Assessment = "ASSESSMENT"
	ADT_Plan       = "PLAN"
)
