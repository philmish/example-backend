package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
    gorm.Model
    id uint `gorm:"primaryKey"`
    email string
    first_name string 
    last_name string
    screen_name string
    is_admin bool
}

type Userdata struct {
    Name string `json:"name"`
    Is_admin bool `json:"isAdmin"`
}

func (u User)ToUserData() *Userdata {
    return &Userdata{Name: u.screen_name, Is_admin: u.is_admin}
}

func UserByEmail(mail, dbname string, user User) (error) {
    db, err := gorm.Open(sqlite.Open(dbname), &gorm.Config{})
    if err != nil {
        return err
    }
    err = db.Where("email = ?", mail).First(&user).Error
    return err
}

