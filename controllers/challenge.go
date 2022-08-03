package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/philmish/example-backend/middleware"
	"github.com/philmish/example-backend/models"
)


func CreateChallenge(c *gin.Context) {
    envVars := c.GetStringMapString("env")
    cookie, err := c.Cookie("token")
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing token cookie"})
        return
    }
    claims, err := middleware.Validate([]byte(envVars["key"]), cookie)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
        return
    }

    if !claims.IsAdmin {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "only admins can create challenges"})
        return
    }

    var req models.ChallengeCreateReq

    decoder := json.NewDecoder(c.Request.Body)
    if err = decoder.Decode(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "malformed request"})
        return
    }

    challenge, err := req.ChallengeFromReq()
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid start date"})
        return
    }

    err = models.CreateChallenge(challenge, envVars["db"])
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, gin.H{"data": "Challenge created successfully"})
}
