package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/nvanonim/fiber-emr/app/configs"
	"github.com/nvanonim/fiber-emr/app/utils"
)

// Protected HandleFunc (middleware) for protected routes
func Protected() gin.HandlerFunc {
	var log = configs.GetLogger()
	return func(c *gin.Context) {
		log.Info("Protected endpoint accessed")
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			log.Error("No token provided")
			c.JSON(http.StatusUnauthorized, utils.GenerateErrorResponse(utils.RC_Unauthorized, utils.RM_Unauthorized))
			c.Abort()
			return
		}

		tokenString = tokenString[7:]
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(utils.GetEnv("JWT_SECRET")), nil
		})

		if err != nil {
			log.Error("Error parsing the token: ", err)
			c.JSON(http.StatusUnauthorized, utils.GenerateErrorResponse(utils.RC_Unauthorized, utils.RM_Unauthorized))
			c.Abort()
			return
		}

		if !token.Valid {
			log.Error("Invalid token")
			c.JSON(http.StatusUnauthorized, utils.GenerateErrorResponse(utils.RC_Unauthorized, utils.RM_Unauthorized))
			c.Abort()
			return
		}

		c.Next()
	}
}
