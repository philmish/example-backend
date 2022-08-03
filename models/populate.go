package models

import (
	"log"
	"time"

	"github.com/philmish/example-backend/middleware"
	"gorm.io/gorm"
)

func populateUsers(db *gorm.DB) {
	hash, err := middleware.Hash("Pass123")
	if err != nil {
		log.Fatalf("Failed to populate users: %s", err.Error())
	}
	user := User{
        E_mail: "test@mail.com",
        First_name: "Jane",
        Last_name: "Doe",
        Screen_name: "JD",
        Pass: hash,
        Is_admin: true
    }
	err = db.Create(&user).Error
	if err != nil {
		log.Fatalf("Failed to populate users: %s", err.Error())
	}
}

