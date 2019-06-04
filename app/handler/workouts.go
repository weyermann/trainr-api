package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/weyermann/trainr-api/app/model"
)

// GetAllWorkouts returns all workouts
func GetAllWorkouts(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	workouts := []model.Workout{}
	db.Find(&workouts)
	respondJSON(w, http.StatusOK, workouts)
}

// CreateWorkout creates a new workout
func CreateWorkout(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	workout := model.Workout{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&workout); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&workout).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, workout)
}

// GetWorkout returns a workout by given workoutID
func GetWorkout(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	workoutID, err := strconv.Atoi(vars["id"])
	if err != nil {
		return
	}
	workout := getWorkoutOr404(db, workoutID, w, r)
	if workout == nil {
		return
	}
	respondJSON(w, http.StatusOK, workout)
}

// UpdateWorkout updates workout data
func UpdateWorkout(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	workoutID, err := strconv.Atoi(vars["id"])
	if err != nil {
		return
	}
	workout := getWorkoutOr404(db, workoutID, w, r)
	if workout == nil {
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&workout); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&workout).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, workout)
}

// DeleteWorkout deletes a workout
func DeleteWorkout(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	workoutID, err := strconv.Atoi(vars["workoutID"])
	if err != nil {
		return
	}
	workout := getWorkoutOr404(db, workoutID, w, r)
	if workout == nil {
		return
	}
	if err := db.Delete(&workout).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusNoContent, nil)
}

// DisableWorkout sets a workout to the passive state
func DisableWorkout(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	workoutID, err := strconv.Atoi(vars["workoutID"])
	if err != nil {
		return
	}
	workout := getWorkoutOr404(db, workoutID, w, r)
	if workout == nil {
		return
	}
	workout.Disable()
	if err := db.Save(&workout).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, workout)
}

// EnableWorkout sets a workout to the active state
func EnableWorkout(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	workoutID, err := strconv.Atoi(vars["workoutID"])
	if err != nil {
		return
	}
	workout := getWorkoutOr404(db, workoutID, w, r)
	if workout == nil {
		return
	}
	workout.Enable()
	if err := db.Save(&workout).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, workout)
}

// getWorkoutOr404 gets a workout instance if exists, or respond the 404 error otherwise
func getWorkoutOr404(db *gorm.DB, workoutID int, w http.ResponseWriter, r *http.Request) *model.Workout {
	workout := model.Workout{}
	if err := db.First(&workout, model.Workout{ID: workoutID}).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &workout
}
