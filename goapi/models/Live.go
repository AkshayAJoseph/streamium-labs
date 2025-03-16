package models

import (
	"gorm.io/gorm"
)

type Live struct {
	LiveID       uint    `gorm:"primaryKey;autoIncrement" json:"id"`
	Title       *string `gorm:"type:varchar(255)" json:"title"`
	Description string  `gorm:"type:text" json:"description"`
	Duration    string  `gorm:"type:varchar(255)" json:"duration"`
	Url 	   string  `gorm:"type:varchar(255)" json:"url"`
}

func MigrateLive(db *gorm.DB) error {
	err := db.AutoMigrate(&Live{})
	return err
}
