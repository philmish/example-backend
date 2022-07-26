package models

import (
    "gorm.io/gorm"
)

type User struct {
    gorm.Model
    id uint `gorm:"primaryKey"`
    email string
    first_name string 
    last_name string
    Screen_name string
    Is_admin bool
}

