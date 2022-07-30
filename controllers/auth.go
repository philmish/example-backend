package controllers

import (
	"net/http"
    "time"

	"github.com/gin-gonic/gin"
	"github.com/philmish/example-backend/middleware"
	"github.com/philmish/example-backend/models"
)

func Auth(c *gin.Context) {
    envVars, err := getEnvVars()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load env"})
    }
    cookie, err := c.Cookie("token")
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "No token"})
        return
    }
    claims, err := midmiddleware.Validate()
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid token"})
        return
    }

    var user models.User
    user, err = models.UserByName(claims["name"], envVars["db"], user)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid token"})
        return
    }

    newClaims := user.ToUserClaims()
    token, err := middleware.CreateToken([]byte(envVars["key"]), newClaims)
    ttl := time.Hour * time.Duration(1)
    now := time.Now()
    expire := now.Add(ttl)
    c.SetCookie("token", token, int(expire.Unix()), "/", "localhost", false, true)
    c.JSON(http.StatusOK, gin.H{"data": "success"})
}
