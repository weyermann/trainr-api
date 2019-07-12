package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Workout struct {
	ID int `gorm:"auto_increment, primary_key" json:"id"`

	// References
	UserID int `json:"userID"` // Foreign key

	Facilities []Facility `gorm:"many2many:workout_facilities" json:"facilities"` // it might make sense to store complete facility objects when the user
	// is allowed to create his own facility types.

	// FacilityIDs []int `gorm:"many2many:workout_facilities" json:"facilities"` // Not working

	WorkoutName       string `json:"workoutName"`
	EnergySystemName  int    `json:"energySystem"`
	EnergySubtypeName int    `json:"energySubtype"`
	Synopsis          string `json:"synopsis"`
	ShortDescription  string `json:"shortDescription"`
	LongDescription   string `json:"longDescription"`
	Duration          int    `json:"duration"`
	ExperienceLevel   int    `json:"experienceLevel"`
	Public            bool   `json:"public"`
	Active            bool   `json:"active"`

	// Default Execution parameters - Initially with these settings, changed when done
	DefNumberOfSets                   int `json:"defNumberOfSets"`
	DefNumberOfRepsPerSet             int `json:"defNumberOfRepsPerSet"`
	DefLoadDurationSeconds            int `json:"defLoadDurationSeconds"`
	DefRestDurationBetweenRepsSeconds int `json:"defRestDurationBetweenRepsSeconds"`
	DefRestDurationBetweenSetsSeconds int `json:"defRestDurationBetweenSetsSeconds"`
}

func (e *Workout) Disable() {
	e.Active = false
}

func (p *Workout) Enable() {
	p.Active = true
}

// DBMigrate will create and migrate the tables, and then make the some relationships if necessary
func DBMigrateWorkout(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&Workout{})
	return db
}
