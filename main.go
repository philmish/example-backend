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
    "github.com/philmish/example-backend/middleware"
)

func main() {
    var populate bool
    var envFile string
    flag.BoolVar(&populate, "populate", false, "if set populate the database")
    flag.StringVar(&envFile, "env", ".env", "env file to use")
    flag.Parse()

    if err := godotenv.Load(envFile); err != nil {
        log.Fatalf(err.Error())
    }

    dbName := os.Getenv("DB_NAME")

    if dbName == "" {
        log.Fatalf("Missing Database name env var")
    }

    if err := models.InitDB(dbName, populate); err != nil {
        log.Fatalf(err.Error())
    }

    r := gin.Default()
    r.Use(middleware.CORS())
    r.Use(middleware.GetEnv())
    
    r.GET("/", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"data": "Hello World"})
    })
    r.GET("/auth", controllers.Auth)
    r.POST("/login", controllers.Login)
    r.POST("/challenge", controllers.CreateChallenge)

    r.Run()
}
