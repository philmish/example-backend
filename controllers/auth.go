package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/philmish/example-backend/middleware"
	"github.com/philmish/example-backend/models"
)

func Auth(c *gin.Context) {
	envVars := c.GetStringMapString("env")

	var user models.User
    claims := c.GetStringMapString("claims")
    user, err := models.UserByName(claims["username"], envVars["db"], user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid token"})
		return
	}

	newClaims := user.ToUserClaims()
	c.Set("claims", newClaims.ToMap())
	middleware.MakeCookie(c)
	c.JSON(http.StatusOK, gin.H{"data": "success"})
}
