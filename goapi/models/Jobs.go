package models

import (
	"gorm.io/gorm"
)

type Job struct {
	JobID       uint    `gorm:"primaryKey;autoIncrement" json:"id"`
	Title       *string `gorm:"type:varchar(255)" json:"title"`
	Description string  `gorm:"type:text" json:"description"`
	Role        string  `gorm:"type:varchar(255)" json:"role"`
	Requirments string  `gorm:"type:text" json:"requirments"`
	PostedOn    string  `gorm:"type:varchar(255)" json:"postedOn"`
}

func MigrateJob(db *gorm.DB) error {
	err := db.AutoMigrate(&Job{})
	return err
}
