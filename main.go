package main

import (
	"log"
	"net/http"
	"os"
    "flag"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/philmish/example-backend/controllers"
	"github.com/philmish/example-backend/models"
)

func main() {
    populate := flag.Bool("populate", false, "if set populate the database")
    envFile := flag.String("env", ".env", "env file to use")

    if err := godotenv.Load(*envFile); err != nil {
        log.Fatalf(err.Error())
    }

    dbName := os.Getenv("DB_NAME")

    if dbName == "" {
        log.Fatalf("Missing Database name env var")
    }

    if err := models.InitDB(dbName, *populate); err != nil {
        log.Fatalf(err.Error())
    }

    r := gin.Default()
    
    r.GET("/", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"data": "Hello World"})
    })

    r.POST("/login", controllers.Login)

    r.Run()
}
