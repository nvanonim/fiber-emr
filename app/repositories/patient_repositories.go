package repositories

import (
	"github.com/nvanonim/fiber-emr/app/models"
	"gorm.io/gorm"
)

func NewPatientRepository(db *gorm.DB) PatientRepository {
	return PatientRepository{db: db}
}

// PatientRepository is the repository for the patient model
type PatientRepository struct {
	db *gorm.DB
}

// Create creates a new patient
func (pr PatientRepository) Create(patient *models.Patient) error {
	if err := pr.db.Create(patient).Error; err != nil {
		return err
	}
	return nil
}

// FindAll finds all patients
func (pr PatientRepository) FindAll() ([]models.Patient, error) {
	var patients []models.Patient
	if err := pr.db.Find(&patients).Error; err != nil {
		return patients, err
	}
	return patients, nil
}

// FindByID finds a patient by id
func (pr PatientRepository) FindByID(id uint) (models.Patient, error) {
	var patient models.Patient
	if err := pr.db.First(&patient, id).Error; err != nil {
		return patient, err
	}
	return patient, nil
}

// FindByMedicalRecordNumber finds a patient by medical record number
func (pr PatientRepository) FindByMedicalRecordNumber(medicalRecordNumber string) (models.Patient, error) {
	var patient models.Patient
	if err := pr.db.Where("medical_record_number = ?", medicalRecordNumber).First(&patient).Error; err != nil {
		return patient, err
	}
	return patient, nil
}
