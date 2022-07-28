package models

import (
	"github.com/philmish/example-backend/middleware"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
    gorm.Model
    E_mail string `gorm:"column:e_mail;type:varchar(255)"`
    First_name string `gorm:"column:first_name;type:varchar(80)"`
    Last_name string `gorm:"column:last_name;type:varchar(80)"`
    Screen_name string `gorm:"column:screen_name;type:varchar(25)"`
    Pass string `gorm:"column:pass;type:varchar(300)"`
    Is_admin bool `gorm:"column:is_admin;type:boolean"` 
}

type Userdata struct {
    Name string `json:"name"`
    Is_admin bool `json:"isAdmin"`
}

func (u User)ToUserData() Userdata {
    return Userdata{Name: u.Screen_name, Is_admin: u.Is_admin}
}

func (u User)ToUserClaims() middleware.UserClaims {
    return middleware.UserClaims{Name: u.Screen_name, IsAdmin: u.Is_admin}
}

func UserByEmail(mail, dbname string, user User) (User, error) {
    db, err := gorm.Open(sqlite.Open(dbname), &gorm.Config{})
    if err != nil {
        return user, err
    }
    err = db.Where(&User{E_mail: mail}).First(&user).Error
    return user, err
}

func UserByName(name, dbname string, user User) (User, error) {
    db, err := gorm.Open(sqlite.Open(dbname), &gorm.Config{})
    if err != nil {
        return user, err
    }
    err = db.Where(&User{Screen_name: name}).First(&user).Error
    return user, err
}

