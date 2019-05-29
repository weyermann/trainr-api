package model

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Session struct {
	gorm.Model
	SessionID             string    `gorm:"unique" json:"sessionID"`
	StartTime             time.Time `json:"start_time"`
	WorkoutExecutionInfos []WorkoutExecutionInfo
}

// DBMigrateSession will create and migrate the tables, and then make the some relationships if necessary
func DBMigrateSession(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&Session{})
	return db
}
