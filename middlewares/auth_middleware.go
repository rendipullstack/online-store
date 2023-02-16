package middlewares

import (
	"net/http"
	"online-store/helpers"
	"online-store/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

func AdminMiddleware(c *gin.Context) {
	role := c.GetString("role")
	if strings.ToLower(role) != "admin" {
		c.AbortWithStatusJSON(
			http.StatusUnauthorized,
			utils.NewErrorResponse(
				http.StatusUnauthorized,
				"Only admin allowed",
				nil,
			),
		)
		return
	}
}

func AuthMiddleware(c *gin.Context) {
	authorizationHeader := c.GetHeader("Authorization")

	if authorizationHeader == "" {
		c.AbortWithStatusJSON(
			http.StatusUnauthorized,
			utils.NewErrorResponse(
				http.StatusUnauthorized,
				"Unauthenticated",
				nil,
			),
		)
		return
	}

	tokenType, token := strings.Split(authorizationHeader, " ")[0], strings.Split(authorizationHeader, " ")[1]
	if tokenType == "" || tokenType != "Bearer" {
		c.AbortWithStatusJSON(
			http.StatusUnauthorized,
			utils.NewErrorResponse(
				http.StatusUnauthorized,
				"Unauthenticated",
				nil,
			),
		)
		return
	}

	if token == "" {
		c.AbortWithStatusJSON(
			http.StatusUnauthorized,
			utils.NewErrorResponse(
				http.StatusUnauthorized,
				"Unauthenticated",
				nil,
			),
		)
		return
	}

	userID, roleUser, err := helpers.ValidateToken(token)
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusUnauthorized,
			utils.NewErrorResponse(
				http.StatusUnauthorized,
				"Unauthenticated",
				err.Error(),
			),
		)
		return
	}

	c.Set("userID", userID)
	c.Set("role", roleUser)

}
