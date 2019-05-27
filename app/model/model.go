package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Workout struct {
	gorm.Model
	WorkoutID   string `gorm:"unique" json:"workoutID"`
	WorkoutName   string `json:"workoutName"`
	EnergySystemName   string `json:"energySystemName"`
	Synopsis   string `json:"synopsis"`
	LongDescription   string `json:"longDescription"`
	Facility   string `json:"facility"`
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
func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&Workout{})
	return db
}
