package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

type Session struct {
	gorm.Model
	SessionID   string `gorm:"unique" json:"sessionID"`
	StartTime   time.Time `json:"start_time"`
	WorkoutExecutionInfos []WorkoutExecutionInfo 

}

// DBMigrate will create and migrate the tables, and then make the some relationships if necessary
func DBMigrateSession(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&Session{})
	return db
}
