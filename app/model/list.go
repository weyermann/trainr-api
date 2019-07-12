package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type ExperienceLevel struct {
	ID   int    `gorm:"auto_increment, primary_key" json:"id"`
	Description string `json:"description"` // Foreign key
}
type EnergySystem struct {
	ID   int    `gorm:"auto_increment, primary_key" json:"id"`
	Description string `json:"description"` // Foreign key
}

// DBMigrateSession will create and migrate the tables, and then make the some relationships if necessary
func DBMigrateLists(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&ExperienceLevel{})
	db.AutoMigrate(&EnergySystem{})
	return db
}
