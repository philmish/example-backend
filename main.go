package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/philmish/example-backend/controllers"
	"github.com/philmish/example-backend/models"
)

func main() {
    if err := models.InitDB("./test.db"); err != nil {
        log.Fatalf(err.Error())
    }

    r := gin.Default()
    
    r.GET("/", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"data": "Hello World"})
    })

    r.POST("/login", controllers.Login)

    r.Run()
}
