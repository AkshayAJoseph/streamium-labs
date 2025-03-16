package models

import (
	"gorm.io/gorm"
)

type Live struct {
	JobID       uint    `gorm:"primaryKey;autoIncrement" json:"id"`
	Title       *string `gorm:"type:varchar(255)" json:"title"`
	Description string  `gorm:"type:text" json:"description"`
	Role        string  `gorm:"type:varchar(255)" json:"role"`
	Requirments string  `gorm:"type:text" json:"requirments"`
	PostedOn    string  `gorm:"type:varchar(255)" json:"postedOn"`
}

func MigrateLive(db *gorm.DB) error {
	err := db.AutoMigrate(&Live{})
	return err
}
