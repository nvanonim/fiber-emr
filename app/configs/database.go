package configs

import (
	"github.com/nvanonim/fiber-emr/app/models"
	"github.com/nvanonim/fiber-emr/app/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

// SetupDB sets up the database
func SetupDB() {
	db = ConnectDB()
	AutoMigrate()
}

// GetDB returns the database
func GetDB() *gorm.DB {
	return db
}

// ConnectDB connects to the database
func ConnectDB() *gorm.DB {
	dsn := utils.GetEnv("DATABASE_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

// AutoMigrate migrates the database
func AutoMigrate() {
	db.AutoMigrate(&models.User{}, &models.Patient{}, &models.AppointmentType{}, &models.Appointment{}, &models.AppointmentData{}, &models.AppointmentDataType{})
}

// CloseDB closes the database
func CloseDB() {
	db, _ := db.DB()
	db.Close()
}
