package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func makeMigrations(db *gorm.DB) (error) {
    err := db.AutoMigrate(&User{})
    if err != nil {
        return err
    }
    err = db.AutoMigrate(&Challenge{})
    return err
}

func populateDb(db *gorm.DB) {
    populateUsers(db)
    populateChallenges(db)
}

func InitDB(db_file string, populate bool) (error) {
    db, err := gorm.Open(sqlite.Open(db_file), &gorm.Config{})

    if err != nil {
        return err
    }

    err = makeMigrations(db)

    if populate {
        populateDb(db)
    }
    return err
}
