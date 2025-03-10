package models

import (
	"gorm.io/gorm"
)

type User struct {
	ID    uint   `gorm:"primaryKey" json:"id"`
	Name  string `json:"name"`
	Email string `gorm:"uniqueIndex" json:"email"`
	Age   int    `json:"age"`
}

func MigrateUserTable(db *gorm.DB) {
	db.AutoMigrate(&User{})
}
