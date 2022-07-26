package middleware

import (
    "golang.org/x/crypto/bcrypt"
)

func Hash(pass string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(pass), 10)
    return string(bytes), err
}

func CheckPass(pass, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass))
    return err == nil
}
