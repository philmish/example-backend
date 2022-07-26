package models

import (
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

var DB *gorm.DB

func makeMigrations(db *gorm.DB) (error) {
    err := db.AutoMigrate(&User{})
    return err
}

func InitDB(db_file string) (error) {
    db, err := gorm.Open(sqlite.Open(db_file), &gorm.Config{})

    if err != nil {
        return err
    }

    err = makeMigrations(db)
    return err
}
