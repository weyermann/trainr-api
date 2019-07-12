package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Facility struct {
	ID          int    `gorm:"auto_increment, primary_key" json:"id"`
	Description string `json:"description"` // Foreign key
}

// DBMigrateFacility will create and migrate the tables, and then make the some relationships if necessary
func DBMigrateFacility(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&Facility{})
	return db
}
