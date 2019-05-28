package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Workout struct {
	gorm.Model
	WorkoutID   string `gorm:"unique" json:"workoutID"`
	WorkoutName   string  `gorm:"unique" json:"workoutName"`
	EnergySystemName   string `json:"energySystemName"`
	EnergySubtypeName string `json:"energySubtypeName"`
	Synopsis   string `json:"synopsis"`
	ShortDescription string `json:"shortDescription"`
	LongDescription   string `json:"longDescription"`
	Facility   string `json:"facility"`
	FacilityOpt   string `json:"facility_opt"`
	Duration    int    `json:"duration"`
	Active 	bool	`json:"active"`
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
