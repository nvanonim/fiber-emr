package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nvanonim/fiber-emr/app/configs"
	"github.com/nvanonim/fiber-emr/app/models"
	"github.com/nvanonim/fiber-emr/app/repositories"
	"github.com/nvanonim/fiber-emr/app/utils"
	"github.com/nvanonim/fiber-emr/pkg/logger"
)

// PatientController is the controller for the patient model
type PatientController struct {
	repo repositories.PatientRepository
	log  *logger.Logger
}

// NewPatientController returns a new patient controller
func NewPatientController(repo repositories.PatientRepository) *PatientController {
	return &PatientController{
		repo: repo,
		log:  configs.GetLogger(),
	}
}

// AddPatient creates a new patient
func (pc PatientController) AddPatient(c *gin.Context) {
	var patient models.PatientRequest

	// Bind the patient struct, if the binding fails return an error
	if err := c.ShouldBindJSON(&patient); err != nil {
		pc.log.Error("Error binding the patient struct: ", err)
		c.JSON(http.StatusBadRequest, utils.GenerateErrorResponse(utils.RC_InvalidRequest, "Invalid request"))
		return
	}

	// Check if the medical record number already exists
	_, err := pc.repo.FindByMedicalRecordNumber(patient.MedicalRecordNumber)
	if err == nil {
		pc.log.Error("Medical record number already exists: ", err)
		c.JSON(http.StatusConflict, utils.GenerateErrorResponse(utils.RC_DataAlreadyExist, "Medical record number already exists"))
		return
	}

	// process data
	// validate the birth date
	birthDate, err := time.Parse(models.DOB, patient.BirthDate)
	if err != nil {
		pc.log.Error("Error parsing the birth date: ", err)
		c.JSON(http.StatusBadRequest, utils.GenerateErrorResponse(utils.RC_InvalidRequest, "Invalid Date Format"))
		return
	}
	// validate the phone number
	phoneNumber := utils.ValidatePhoneNumber(patient.PhoneNumber)
	pc.log.Debug("Formatted Phone Number: ", phoneNumber)
	if phoneNumber == "" {
		pc.log.Error("Error validating the phone number: ", err)
		c.JSON(http.StatusBadRequest, utils.GenerateErrorResponse(utils.RC_InvalidRequest, "Invalid Phone Number"))
		return
	}

	// Create the patient
	patientModel := models.Patient{
		MedicalRecordNumber: patient.MedicalRecordNumber,
		Name:                patient.Name,
		Gender:              patient.Gender,
		BirthDate:           birthDate,
		Address:             patient.Address,
		PhoneNumber:         phoneNumber,
	}

	if err := pc.repo.Create(&patientModel); err != nil {
		pc.log.Error("Error creating the patient: ", err)
		c.JSON(http.StatusInternalServerError, utils.GenerateErrorResponse(utils.RC_InternalServerError, "Error creating the patient"))
		return
	}

	pc.log.Infof("Patient created for medical record number: %s", patientModel.MedicalRecordNumber)

	c.JSON(http.StatusCreated, utils.GenerateResponse(utils.RC_Success, utils.RM_Success, gin.H{"id": patientModel.ID}))
}

// ListPatients returns all patients
func (pc PatientController) ListPatients(c *gin.Context) {
	patients, err := pc.repo.FindAll()
	if err != nil {
		pc.log.Error("Error finding all patients: ", err)
		c.JSON(http.StatusInternalServerError, utils.GenerateErrorResponse(utils.RC_InternalServerError, "Error finding all patients"))
		return
	}

	// map the patients to the response
	var response []models.PatientResponse
	for _, patient := range patients {
		response = append(response, models.PatientResponse{
			ID:                  patient.ID,
			MedicalRecordNumber: patient.MedicalRecordNumber,
			Name:                patient.Name,
			Gender:              patient.Gender,
			BirthDate:           patient.BirthDate.Format(models.DOB),
			Address:             patient.Address,
			PhoneNumber:         patient.PhoneNumber,
		})
	}
	c.JSON(http.StatusOK, utils.GenerateResponse(utils.RC_Success, utils.RM_Success, response))
}

// GetPatient returns a patient by id
func (pc PatientController) GetPatient(c *gin.Context) {
	id, err := utils.StringToUint(c.Param("id"))
	if err != nil {
		pc.log.Error("Error converting the id to uint: ", err)
		c.JSON(http.StatusBadRequest, utils.GenerateErrorResponse(utils.RC_InvalidRequest, "Invalid request"))
		return
	}
	patient, err := pc.repo.FindByID(id)
	if err != nil {
		pc.log.Error("Error finding the patient by id: ", err)
		c.JSON(http.StatusInternalServerError, utils.GenerateErrorResponse(utils.RC_InternalServerError, "Patient not found"))
		return
	}

	// map the patient to the response
	patientResponse := models.PatientResponse{
		ID:                  patient.ID,
		MedicalRecordNumber: patient.MedicalRecordNumber,
		Name:                patient.Name,
		Gender:              patient.Gender,
		BirthDate:           patient.BirthDate.Format("2006-01-02"),
		Address:             patient.Address,
		PhoneNumber:         patient.PhoneNumber,
	}

	c.JSON(http.StatusOK, utils.GenerateResponse(utils.RC_Success, utils.RM_Success, patientResponse))
}
