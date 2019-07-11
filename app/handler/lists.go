package handler

import (
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/weyermann/trainr-api/app/model"
)

// GetAllWorkouts returns all workouts
func GetAllFacilities(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	facilities := []model.Facility{}
	db.Find(&facilities)
	respondJSON(w, http.StatusOK, facilities)
}