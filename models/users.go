package models

import (
    "gorm.io/gorm"
    "gorm.io/driver/sqlite"
)

type User struct {
    gorm.Model
    id uint `json:"id" gorm:"primaryKey"`
    email string `json:"email"`
    first_name string `json:"first_name"`
    last_name string `json:"last_name"`
    screen_name string `json:"screen_name"`
    is_admin bool `json:"is_admin"`
}

type LoginResponse struct {
    name string `json:"name"`
    is_admin bool `json:"is_admin"`
}

