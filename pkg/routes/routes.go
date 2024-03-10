package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nvanonim/fiber-emr/pkg/controllers"
	"github.com/nvanonim/fiber-emr/pkg/utils"
)

// register the routes
func RegisterRoutes(r *gin.Engine) {
	// ping
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	AuthRoutes(r)
}

func AuthRoutes(r *gin.Engine) {
	auth := r.Group("/auth")
	// auth
	auth.POST("/login", controllers.Login)
	auth.POST("/signup", controllers.Signup)
	auth.GET("/validate", controllers.Protected(), func(c *gin.Context) {
		c.JSON(http.StatusOK, utils.GenerateResponse(utils.RC_Success, utils.RM_Success))
	})
}

// func PatientRoutes(r *gin.Engine) {
// 	patient := r.Group("/patient")
// 	patient.POST("/create", controllers.CreatePatient)
// 	patient.GET("/list", controllers.ListPatients)
// 	patient.GET("/get/:id", controllers.GetPatient)
// 	patient.PUT("/update/:id", controllers.UpdatePatient)
// 	patient.DELETE("/delete/:id", controllers.DeletePatient)
// }
