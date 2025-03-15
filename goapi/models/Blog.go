package models

import (
	"gorm.io/gorm"
)

type Blog struct {
	BlogID       uint    `gorm:"primaryKey;autoIncrement" json:"id"`
	Title        *string `gorm:"type:varchar(255)" json:"title"`
	Content      string  `gorm:"type:text" json:"content"`
	Banner       string  `gorm:"type:varchar(255)" json:"banner"`
	Date         string  `gorm:"type:varchar(255)" json:"date"`
	PostedByName string  `gorm:"type:varchar(255)" json:"postedName"`
	PostedByID   uint    `gorm:"type:int" json:"postedId"`
}

func MigrateBlog(db *gorm.DB) error {
	err := db.AutoMigrate(&User{})
	return err
}
