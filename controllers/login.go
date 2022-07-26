package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	//"github.com/philmish/example-backend/middleware"
	"github.com/philmish/example-backend/models"
)

type LoginRequest struct {
    Email string `json:"email"`
    Pass string `json:"password"`
}

type LoginResponse struct {
    Name string `json:"name"`
    Is_admin bool `json:"isAdmin"`
}

func Login(c *gin.Context) {
    var req LoginRequest
    var user models.User

    if err := c.BindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
    }

    if err := models.DB.Where("email = ?", req.Email).First(&user).Error; err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Incorrect creds"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": LoginResponse{Name: user.Screen_name, Is_admin: user.Is_admin}})
}
