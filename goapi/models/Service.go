package models

import (
	"gorm.io/gorm"
)

type S	ervice struct {
	ServiceID       uint    `gorm:"primaryKey;autoIncrement" json:"id"`
	Type 		  *string `gorm:"type:varchar(255)" json:"type"`
	Category	  string  `gorm:"type:varchar(255)" json:"category"`
	Status		  string  `gorm:"type:varchar(255)" json:"status"`
	Price		  uint    `gorm:"type:int" json:"price"`
}

func MigrateService(db *gorm.DB) error {
	err := db.AutoMigrate(&Service{})
	return err
}
