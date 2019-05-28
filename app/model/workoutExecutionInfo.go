package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type WorkoutExecutionInfo struct {
	gorm.Model
	// Define properties
	WorkoutID int
	Workout Workout
}

// DBMigrate will create and migrate the tables, and then make the some relationships if necessary
func DBMigrateWorkoutExecutionInfo(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&WorkoutExecutionInfo{})
	return db
}
