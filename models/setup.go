package models

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func makeMigrations(db *gorm.DB) (error) {
    err := db.AutoMigrate(&User{})
    return err
}

func populateDb() {
    //TODO implement data population
    fmt.Println("Populating db...")
}

func InitDB(db_file string, populate bool) (error) {
    db, err := gorm.Open(sqlite.Open(db_file), &gorm.Config{})

    if err != nil {
        return err
    }

    if populate {
        populateDb()
    }

    err = makeMigrations(db)
    return err
}
