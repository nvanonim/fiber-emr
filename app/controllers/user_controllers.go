package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nvanonim/fiber-emr/app/configs"
	"github.com/nvanonim/fiber-emr/app/models"
	"github.com/nvanonim/fiber-emr/app/repositories"
	"github.com/nvanonim/fiber-emr/app/utils"
	"github.com/nvanonim/fiber-emr/pkg/logger"
	"golang.org/x/crypto/bcrypt"
)

// UserController is the controller for user-related operations
type UserController struct {
	UserRepository repositories.UserRepository
	log            *logger.Logger
}

// NewUserController creates a new UserController
func NewUserController(repo repositories.UserRepository) *UserController {
	return &UserController{
		UserRepository: repo,
		log:            configs.GetLogger(),
	}
}

// Signup creates a new user
func (uc *UserController) Signup(c *gin.Context) {
	var userRegister models.UserRegister

	if err := c.ShouldBindJSON(&userRegister); err != nil {
		uc.log.Error("Error binding the user register struct: ", err)
		c.JSON(http.StatusBadRequest, utils.GenerateErrorResponse(utils.RC_InvalidRequest, "Invalid request"))
		return
	}

	dbUser, _ := uc.UserRepository.FindByUsername(userRegister.Username)

	if dbUser.ID != 0 {
		c.JSON(http.StatusConflict, utils.GenerateErrorResponse(utils.RC_DataAlreadyExist, "Username already taken"))
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(userRegister.Password), bcrypt.DefaultCost)
	if err != nil {
		uc.log.Error("Error hashing the password: ", err)
		c.JSON(http.StatusInternalServerError, utils.GenerateErrorResponse(utils.RC_InternalServerError, "Error hashing the password"))
		return
	}

	userRegister.Password = string(hash)

	user := models.User{
		Username: userRegister.Username,
		Name:     userRegister.Name,
		Password: userRegister.Password,
	}

	if err := uc.UserRepository.Create(&user); err != nil {
		uc.log.Error("Error creating the user: ", err)
		c.JSON(http.StatusInternalServerError, utils.GenerateErrorResponse(utils.RC_InternalServerError, "Error creating the user"))
		return
	}

	uc.log.Infof("User created for username: %s", user.Username)

	c.JSON(http.StatusOK, utils.GenerateResponse(utils.RC_Success, "User created successfully"))
}

// Login logs in a user
func (uc *UserController) Login(c *gin.Context) {
	var user models.UserLogin

	if err := c.ShouldBindJSON(&user); err != nil {
		uc.log.Error("Error binding the user login struct: ", err)
		c.JSON(http.StatusBadRequest, utils.GenerateErrorResponse(utils.RC_InvalidRequest, "Invalid request"))
		return
	}

	dbUser, err := uc.UserRepository.FindByUsername(user.Username)
	if err != nil {
		uc.log.Error("Error finding the user by username: ", err)
		c.JSON(http.StatusInternalServerError, utils.GenerateErrorResponse(utils.RC_DataNotFound, "User not found"))
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password)); err != nil {
		uc.log.Error("Invalid password: ", err)
		c.JSON(http.StatusUnauthorized, utils.GenerateErrorResponse(utils.RC_Unauthorized, "Invalid password"))
		return
	}

	token, err := utils.CreateToken(dbUser.ID)
	if err != nil {
		uc.log.Error("Error creating the token:", err)
		c.JSON(http.StatusInternalServerError, utils.GenerateErrorResponse(utils.RC_InternalServerError, "Error creating token"))
		return
	}

	response := models.UserResponse{
		Username:       dbUser.Username,
		Name:           dbUser.Name,
		Token:          token,
		ExpirationTime: int64(utils.JWTExpirationTime.Seconds()),
	}

	uc.log.Debugf("User logged in for username: %s", dbUser.Username)

	c.JSON(http.StatusOK, utils.GenerateResponse(utils.RC_Success, "Login successful", response))
}
