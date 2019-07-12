package handler

import (
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/weyermann/trainr-api/app/model"
)

// GetAllFacilities returns all facilities
func GetAllFacilities(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	facilities := []model.Facility{}
	db.Find(&facilities)
	respondJSON(w, http.StatusOK, facilities)
}

// GetAllEnergySystems returns all energy systems
func GetAllEnergySystems(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	energysystems := []model.EnergySystem{}
	db.Find(&energysystems)
	respondJSON(w, http.StatusOK, energysystems)
}

// GetAllExperienceLevels returns all energy systems
func GetAllExperienceLevels(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	experiencelevels := []model.ExperienceLevel{}
	db.Find(&experiencelevels)
	respondJSON(w, http.StatusOK, experiencelevels)
}