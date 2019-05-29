package model

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Session struct {
	ID int `gorm:"auto_increment, primary_key" json:"sessionID"`

	// References
	UserID int `json:"workoutID"` // Foreign key
	// User User // the referenced user - Session belongs to exactly one user

	StartTime             time.Time `json:"start_time"`
	WorkoutExecutionInfos []WorkoutExecution
}

// DBMigrateSession will create and migrate the tables, and then make the some relationships if necessary
func DBMigrateSession(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&Session{})
	return db
}
