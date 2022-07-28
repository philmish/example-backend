package middleware

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
)

type UserClaims struct {
    Name string
    IsAdmin bool
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
