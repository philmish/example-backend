package middleware

import (
	"fmt"
	"net/http"
    "time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type UserClaims struct {
    Name string
    IsAdmin bool
}

func (u UserClaims)ToMap() map[string]string {
    var admin string
    if u.IsAdmin {
        admin = "yes"
    } else {
        admin = "no"
    }
    return map[string]string{
        "username": u.Name,
        "isAdmin": admin,
    }
}

func fromMap(claims map[string]string) UserClaims {
    var admin bool
    if claims["isAdmin"] == "yes" {
        admin = true
    } else {
        admin = false
    }
    return UserClaims{Name: claims["username"], IsAdmin: admin}
} 

func Validate(key []byte, tokenString string) (UserClaims, error) {
    token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
        if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("Invalid signing Method: %v", t.Header["alg"])
        }
        return key, nil
    })

    if err != nil {
        return UserClaims{}, err
    }

    if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
        return UserClaims{Name: claims["name"].(string), IsAdmin: claims["isadmin"].(bool)}, nil
    } else {
        return UserClaims{}, fmt.Errorf("Invalid claims")
    }
}

func CreateToken(key []byte, claims UserClaims) (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "name": claims.Name,
        "isadmin": claims.IsAdmin,
    })

    tokenString, err := token.SignedString(key)
    return tokenString, err
}

func ParseToken() gin.HandlerFunc {
    return func(c *gin.Context) {
        env := c.GetStringMapString("env")
        cookie, err := c.Cookie("token")
        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "No token"})
            c.Abort()
            return
        }
        claims, err := Validate([]byte(env["key"]), cookie)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid token"})
            return
        }
        c.Set("claims", claims.ToMap())
        c.Next()
    }
}

func MakeCookie(c *gin.Context) {
    env := c.GetStringMapString("env")
    claimsMap := c.GetStringMapString("claims")
    claims := fromMap(claimsMap)
    token, err := CreateToken([]byte(env["key"]), claims)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        c.Abort()
        return
    }
    ttl := time.Hour * time.Duration(1)
    now := time.Now()
    expire := now.Add(ttl)
    c.SetCookie("token", token, int(expire.Unix()), "/", "localhost", false, true)
    return
}
