package middleware

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func getEnvVars() (map[string]string, error) {
	res := map[string]string{}
	res["key"] = os.Getenv("SECRET")
	res["db"] = os.Getenv("DB_NAME")
	for k, v := range res {
		if v == "" {
			return res, fmt.Errorf("No value for %s found", k)
		}
	}
	return res, nil
}

func GetEnv() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		envVars, err := getEnvVars()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			ctx.Abort()
			return
		}
		ctx.Set("env", envVars)
		ctx.Next()
	}
}
