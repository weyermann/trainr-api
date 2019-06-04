package handler

import (
	"encoding/json"
	// "fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/weyermann/trainr-api/app/model"
)

// GetAllExecutions returns all sessions of a user
func GetAllExecutions(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	executions := []model.WorkoutExecution{}

	vars := mux.Vars(r)
	sessionID := vars["sessionID"]

	// Get all sessions matching the user
	db.Where("sessionID = ?", sessionID).Find(&executions)
	//// SELECT * FROM users WHERE sessionID = 'xyz';

	respondJSON(w, http.StatusOK, executions)
}

// GetExecution returns a workout by given executionID
func GetExecution(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	executionID, err := strconv.Atoi(vars["executionID"])
	if err != nil {
		return
	}
	execution := getExecutionOr404(db, executionID, w, r)
	if execution == nil {
		return
	}
	respondJSON(w, http.StatusOK, execution)
}

// CreateWorkout creates a new workout
func CreateExecution(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	execution := model.WorkoutExecution{}
	// workout := model.Workout{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&execution); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	// Read from workout Table where ID = execution.WorkoutID
	db.First(&execution.Workout, execution.WorkoutID)

	if err := db.Save(&execution).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, execution)
}

// getExecutionOr404 gets a workout execution instance if exists, or respond the 404 error otherwise
func getExecutionOr404(db *gorm.DB, executionID int, w http.ResponseWriter, r *http.Request) *model.WorkoutExecution {
	execution := model.WorkoutExecution{}
	if err := db.First(&execution, model.WorkoutExecution{ID: executionID}).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &execution
}
