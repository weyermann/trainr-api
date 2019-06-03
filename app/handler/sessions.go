package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/weyermann/trainr-api/app/model"
)

// GetAllSessions returns all sessions of a user
func GetAllSessions(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	sessions := []model.Session{}

	vars := mux.Vars(r)
	userID := vars["userID"]

	// Get all sessions matching the user
	db.Where("userID = ?", userID).Find(&sessions)
	//// SELECT * FROM users WHERE name = 'jinzhu';

	respondJSON(w, http.StatusOK, sessions)
}

// GetSession returns a workout by given workoutID
func GetSession(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	sessionID, err := strconv.Atoi(vars["sessionID"])
	if err != nil {
		return
	}
	session := getSessionOr404(db, sessionID, w, r)
	if session == nil {
		return
	}
	respondJSON(w, http.StatusOK, session)
}

// CreateSession creates a new session instance
func CreateSession(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	session := model.Session{}

	decoder := json.NewDecoder(r.Body)

	// decoder needs to convert the startTime string into a time object, via parse 
	// somehow like so:
	/* 
	t, err := time.Parse(time.RFC3339, "2006-01-02T15:04:05-07:00")
	if err != nil {
		log.Fatal(err)
	} */
	if err := decoder.Decode(&session); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&session).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, session)
}

// UpdateSession updates session data (e.g. move to another day)
func UpdateSession(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	sessionID, err := strconv.Atoi(vars["sessionID"])
	if err != nil {
		return
	}
	session := getSessionOr404(db, sessionID, w, r)
	if session == nil {
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&session); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&session).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, session)
}

// DeleteSession deletes a session
func DeleteSession(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	sessionID, err := strconv.Atoi(vars["sessionID"])
	if err != nil {
		return
	}
	session := getSessionOr404(db, sessionID, w, r)
	if session == nil {
		return
	}
	// TODO add on Delete cascade
	if err := db.Delete(&session).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusNoContent, nil)
}

// getSessionOr404 gets a session instance if exists, or respond the 404 error otherwise
func getSessionOr404(db *gorm.DB, sessionID int, w http.ResponseWriter, r *http.Request) *model.Session {
	session := model.Session{}
	if err := db.First(&session, model.Session{ID: sessionID}).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &session
}
