package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type WorkoutExecution struct {
	ID int `gorm:"auto_increment, primary_key" json:"id"`
	// References
	// WorkoutExecutionInfo belongs to one workout
	WorkoutID int `json:"workoutID"` // Foreign key

	/* The referenced workout: WorkoutExecution belongs to exactly one workout
	When this is incommented, a workout object will be sent with the payload
	*/
	Workout Workout

	SessionID int `json:"sessionID"` // ExecutionInfo references exactly one session. A session can have many workoutExecutions

	// Execution parameters - Initially with the workout default settings, changed when done
	NumberOfSets                   int `json:"numberOfSets"`
	NumberOfRepsPerSet             int `json:"numberOfRepsPerSet"`
	LoadDurationSeconds            int `json:"loadDurationSeconds"`
	RestDurationBetweenRepsSeconds int `json:"restDurationBetweenRepsSeconds"`
	RestDurationBetweenSetsSeconds int `json:"restDurationBetweenSetsSeconds"`

	// Log personal info
	IsFinished               bool   `json:"isFinished"`
	ExecutedRatioPercent     int    `json:"executedRatioPercent"` // Maybe this will be calculated after
	ExhaustionLevelOneToFive int    `json:"exhaustionLevelOneToFive"`
	PersonalRemarks          string `json:"personalRemarks"`
	LessonLearned            string `json:"lessonLearned"`
}

// DBMigrateWorkoutExecution will create and migrate the tables, and then make the some relationships if necessary
func DBMigrateWorkoutExecution(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&WorkoutExecution{})
	return db
}
