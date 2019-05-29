package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type WorkoutExecutionInfo struct {
	gorm.Model
	// Define properties
	WorkoutID int `json:"workoutID"` // Foreign key
	Workout Workout // the referenced workout - ExecutionInfo belongs to exactly one workout

	// Execution parameters - Initially with the workout default settings, changed when done
	NumberOfSets int `json:"numberOfSets"`
	NumberOfRepsPerSet int `json:"numberOfRepsPerSet"`
	LoadDurationMins int `json:"loadDurationMins"`
	RestDurationBetweenReps int `json:"restDurationBetweenReps"`
	RestDurationBetweenSets int `json:"restDurationBetweenSets"`

	// Log personal info
	ExecutedRatioPercent int `json:"executedRatioPercent"`
	ExhaustionLevelOneToFive int `json:"exhaustionLevelOneToFive"`
	PersonalRemarks string `json:"personalRemarks"`
	LessonLearned string `json:"lessonLearned"`
}

// DBMigrateWorkoutExecutionInfo will create and migrate the tables, and then make the some relationships if necessary
func DBMigrateWorkoutExecutionInfo(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&WorkoutExecutionInfo{})
	return db
}
