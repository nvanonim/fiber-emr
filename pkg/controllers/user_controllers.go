package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/nvanonim/fiber-emr/pkg/configs"
	"github.com/nvanonim/fiber-emr/pkg/configs/logger"
	"github.com/nvanonim/fiber-emr/pkg/models"
	"github.com/nvanonim/fiber-emr/pkg/repositories"
	"github.com/nvanonim/fiber-emr/pkg/utils"
	"golang.org/x/crypto/bcrypt"
)

var UserRepository = repositories.NewUserRepository()
var Logger = logger.New()

// Signup
func Signup(c *gin.Context) {
	// Get the userRegister register struct (dto)
	var userRegister models.UserRegister

	// Bind the user register struct, if the binding fails return an error
	if err := c.ShouldBindJSON(&userRegister); err != nil {
		Logger.Error("Error binding the user register struct: ", err)
		c.JSON(http.StatusBadRequest, utils.GenerateErrorResponse(utils.RC_InvalidRequest, "Invalid request"))
		return
	}

	// check if username is already taken
	dbUser, err := UserRepository.FindByUsername(userRegister.Username)
	if err != nil {
		Logger.Error("Error finding the user by username: ", err)
		c.JSON(http.StatusInternalServerError, utils.GenerateErrorResponse(utils.RC_InternalServerError, "Error finding the user by username"))
		return
	}

	if dbUser.ID != 0 {
		c.JSON(http.StatusConflict, utils.GenerateErrorResponse(utils.RC_DataAlreadyExist, "Username already taken"))
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(userRegister.Password), bcrypt.DefaultCost)
	if err != nil {
		Logger.Error("Error hashing the password: ", err)
		c.JSON(http.StatusInternalServerError, utils.GenerateErrorResponse(utils.RC_InternalServerError, "Error hashing the password"))
		return
	}

	userRegister.Password = string(hash)

	// map userRegister to user
	user := models.User{
		Username: userRegister.Username,
		Name:     userRegister.Name,
		Password: userRegister.Password,
	}

	if err := UserRepository.Create(&user); err != nil {
		Logger.Error("Error creating the user: ", err)
		c.JSON(http.StatusInternalServerError, utils.GenerateErrorResponse(utils.RC_InternalServerError, "Error creating the user"))
		return
	}

	// Log the successful user creation
	Logger.Infof("User created for username: %s", user.Username)

	c.JSON(http.StatusOK, utils.GenerateResponse(utils.RC_Success, "User created successfully"))
}

// Login
func Login(c *gin.Context) {
	// Get the user login struct (dto)
	var user models.UserLogin

	// Bind the user login struct, if the binding fails return an error
	if err := c.ShouldBindJSON(&user); err != nil {
		Logger.Error("Error binding the user login struct: ", err)
		c.JSON(http.StatusBadRequest, utils.GenerateErrorResponse(utils.RC_InvalidRequest, "Invalid request"))
		return
	}

	db := configs.GetDB()
	var dbUser models.User
	db.Where("username = ?", user.Username).First(&dbUser)

	if dbUser.ID == 0 {
		c.JSON(http.StatusNotFound, utils.GenerateErrorResponse(utils.RC_DataNotFound, "User not found"))
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, utils.GenerateErrorResponse(utils.RC_Unauthorized, "Invalid password"))
		return
	}

	token, err := utils.CreateToken(dbUser.ID)
	if err != nil {
		Logger.Error("Error creating the token:", err)
		c.JSON(http.StatusInternalServerError, utils.GenerateErrorResponse(utils.RC_InternalServerError, "Error creating token"))
		return
	}

	response := models.UserResponse{
		Username: dbUser.Username,
		Name:     dbUser.Name,
		Token:    token,
		// Expiration time in seconds
		ExpirationTime: int64(utils.JWTExpirationTime.Seconds()),
	}

	Logger.Debugf("User logged in for username: %s", dbUser.Username)

	c.JSON(http.StatusOK, utils.GenerateResponse(utils.RC_Success, "Login successful", response))
}

// Protected HandleFunc
func Protected() gin.HandlerFunc {
	return func(c *gin.Context) {
		Logger.Info("Protected endpoint accessed")
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			Logger.Error("No token provided")
			c.JSON(http.StatusUnauthorized, utils.GenerateErrorResponse(utils.RC_Unauthorized, utils.RM_Unauthorized))
			c.Abort()
			return
		}

		tokenString = tokenString[7:]
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(utils.GetEnv("JWT_SECRET")), nil
		})

		if err != nil {
			Logger.Error("Error parsing the token: ", err)
			c.JSON(http.StatusUnauthorized, utils.GenerateErrorResponse(utils.RC_Unauthorized, utils.RM_Unauthorized))
			c.Abort()
			return
		}

		if !token.Valid {
			Logger.Error("Invalid token")
			c.JSON(http.StatusUnauthorized, utils.GenerateErrorResponse(utils.RC_Unauthorized, utils.RM_Unauthorized))
			c.Abort()
			return
		}

		c.Next()
	}
}
