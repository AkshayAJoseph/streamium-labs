package models

import (
	"gorm.io/gorm"
)

type Esport struct {
	EsportID   uint    `gorm:"primaryKey;autoIncrement" json:"id"`
	Game       *string `gorm:"type:varchar(255)" json:"game"`
	Team1      string  `gorm:"type:varchar(255)" json:"team1"`
	Team2      string  `gorm:"type:varchar(255)" json:"team2"`
	Team1Score uint    `gorm:"type:int" json:"team1Score"`
	Team2Score uint    `gorm:"type:int" json:"team2Score"`
	Team1Logo  string  `gorm:"type:varchar(255)" json:"team1Logo"`
	Team2Logo  string  `gorm:"type:varchar(255)" json:"team2Logo"`
	Fee        uint    `gorm:"type:int" json:"fee"`
}

func MigrateEsport(db *gorm.DB) error {
	err := db.AutoMigrate(&Esport{})
	return err
}
