package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nvanonim/fiber-emr/app/configs"
	"github.com/nvanonim/fiber-emr/app/controllers"
	"github.com/nvanonim/fiber-emr/app/middlewares"
	"github.com/nvanonim/fiber-emr/app/repositories"
	"github.com/nvanonim/fiber-emr/app/utils"
	"gorm.io/gorm"
)

// register the routes
func RegisterRoutes(r *gin.Engine) {
	db := configs.GetDB()

	// ping
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	AuthRoutes(db, r)
	PatientRoutes(db, r)
}

func AuthRoutes(db *gorm.DB, r *gin.Engine) {
	userRepo := repositories.NewUserRepository(db)
	userController := controllers.NewUserController(userRepo)

	auth := r.Group("/auth")
	// auth
	auth.POST("/login", userController.Login)
	auth.POST("/signup", userController.Signup)
	auth.GET("/validate", middlewares.Protected(), func(c *gin.Context) {
		c.JSON(http.StatusOK, utils.GenerateResponse(utils.RC_Success, utils.RM_Success))
	})
}

func PatientRoutes(db *gorm.DB, r *gin.Engine) {
	patientRepo := repositories.NewPatientRepository(db)
	patientController := controllers.NewPatientController(patientRepo)

	patient := r.Group("/patient", middlewares.Protected())
	patient.POST("/add", patientController.AddPatient)
	patient.GET("/list", patientController.ListPatients)
	patient.GET("/get/:id", patientController.GetPatient)
}
