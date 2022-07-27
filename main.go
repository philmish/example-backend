package main

import (
	"log"
	"net/http"
	"os"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/philmish/example-backend/controllers"
	"github.com/philmish/example-backend/models"
)

func main() {

    if err := godotenv.Load(".env"); err != nil {
        log.Fatalf(err.Error())
    }

    dbName := os.Getenv("DB_NAME")

    if dbName == "" {
        log.Fatalf("Missing Database name env var")
    }

    if err := models.InitDB(dbName); err != nil {
        log.Fatalf(err.Error())
    }

    r := gin.Default()
    
    r.GET("/", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"data": "Hello World"})
    })

    r.POST("/login", controllers.Login)

    r.Run()
}
