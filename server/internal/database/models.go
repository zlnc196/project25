package database

import (
	"log"

	_ "github.com/lib/pq"
	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique"`
	Password string 
}

func CreateUserDB(db *gorm.DB, username string, password string) {
	newRow := User{Username:username, Password:password}
	err := db.Model(&User{}).Create(&newRow).Error
	if err != nil {
		log.Printf("Error in creating new user due to: %s", err)
	}
}