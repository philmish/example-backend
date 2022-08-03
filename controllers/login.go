package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/philmish/example-backend/middleware"
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
    envVars := c.GetStringMapString("env")
    var req LoginRequest
    var user models.User

    decoder := json.NewDecoder(c.Request.Body)
    if err := decoder.Decode(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "malformed data"})
    }

    user, err := models.UserByEmail(req.Email, envVars["db"], user)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid creds"})
        return
    }

    if !middleware.CheckPass(req.Pass, user.Pass) {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid creds"})
        return
    }

    data := user.ToUserData()
    claims := user.ToUserClaims() 
    token, err := middleware.CreateToken([]byte(envVars["key"]), claims)
    ttl := time.Hour * time.Duration(1)
    now := time.Now()
    expire := now.Add(ttl)
    c.SetCookie("token", token, int(expire.Unix()), "/", "localhost", false, true)

    c.JSON(http.StatusOK, gin.H{"data": data, "token": token})
}
