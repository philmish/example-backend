package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Logout(c *gin.Context) {
	c.SetCookie("token", "", 0, "/", "localhost", false, true)
    c.JSON(http.StatusOK, gin.H{"data": "success"})
}
