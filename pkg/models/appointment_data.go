package models

import (
	"time"

	"github.com/jackc/pgtype"
	"gorm.io/gorm"
)

// Appointment Data Model with created_at and updated_at fields. Also having images, pdf in jsonb format to store multiple images and pdf

type AppointmentData struct {
	gorm.Model
	ID             uint         `json:"id" gorm:"primaryKey"` // gorm:"primaryKey" is used to set the primary key
	AppointmentID  uint         `json:"appointment_id"`
	DataTypeID     uint         `json:"data_type_id"`
	Images         pgtype.JSONB `json:"images"`
	Pdf            pgtype.JSONB `json:"pdf"`
	AdditionalInfo string       `json:"additional_info"`
	CreatedAt      time.Time    `json:"-"`
	UpdatedAt      time.Time    `json:"-"`
}

// AppointmentDataResponse is the response for the appointment data model
type AppointmentDataResponse struct {
	ID             uint   `json:"id"`
	AppointmentID  uint   `json:"appointment_id"`
	DataTypeID     uint   `json:"data_type_id"`
	Images         string `json:"images"`
	Pdf            string `json:"pdf"`
	AdditionalInfo string `json:"additional_info"`
}
