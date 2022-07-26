package models

import (
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

func makeMigrations(db *gorm.DB) (error) {
    err := db.AutoMigrate(&User{}); err != nil {
    return err
}

func GetDB(db_file string) (*gorm.DB, error) {
    db, err := gorm.Open(db_file, &gorm.Config{})

    if err != nil {
        return nil, err
    }

    err = makeMigrations(db)
    if err != nil {
        return nil, err
    }

    return db, err
}
