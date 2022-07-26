package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/philmish/example-backend/models"
    "github.com/philmish/example-backend/controllers"
)

func main() {
    models.InitDB("./test.db")

    r := gin.Default()
    
    r.GET("/", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"data": "Hello World"})
    })

    r.POST("/login", controllers.Login)

    r.Run()
}
