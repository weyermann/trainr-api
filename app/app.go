package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/weyermann/trainr-api/app/handler"
	"github.com/weyermann/trainr-api/app/model"
	"github.com/weyermann/trainr-api/config"
)

// App has router and db instances
type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

// App initialize with predefined configuration
func (a *App) Initialize(config *config.Config) {
	dbURI := fmt.Sprintf("%s:%s@/%s?charset=%s&parseTime=True",
		config.DB.Username,
		config.DB.Password,
		config.DB.Name,
		config.DB.Charset)

	db, err := gorm.Open(config.DB.Dialect, dbURI)
	if err != nil {
		log.Fatal("Could not connect database")
	}

	a.DB = model.DBMigrateWorkout(db)
	a.DB = model.DBMigrateUser(db)
	a.DB = model.DBMigrateSession(db)
	a.DB = model.DBMigrateWorkoutExecution(db)
	a.Router = mux.NewRouter()
	a.setRouters()
}

// Set all required routers
func (a *App) setRouters() {
	// Workouts
	a.Get("/workouts", a.GetAllWorkouts)
	a.Post("/workouts", a.CreateWorkout)
	a.Get("/workouts/{id}", a.GetWorkout)
	a.Put("/workouts/{id}", a.UpdateWorkout)
	a.Delete("/workouts/{id}", a.DeleteWorkout)
	a.Put("/workouts/{id}/disable", a.DisableWorkout)
	a.Put("/workouts/{id}/enable", a.EnableWorkout)

	// Sessions
	a.Get("/sessions", a.GetAllSessions)
	a.Post("/sessions", a.CreateSession)
	a.Get("/sessions/{id}", a.GetSession)
	a.Put("/sessions/{id}", a.UpdateSession)
	a.Delete("/sessions/{id}", a.DeleteSession)
}

// Wrap the router for GET method
func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}

// Wrap the router for POST method
func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}

// Wrap the router for PUT method
func (a *App) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("PUT")
}

// Wrap the router for DELETE method
func (a *App) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("DELETE")
}

// Handlers to manage Workout Data
func (a *App) GetAllWorkouts(w http.ResponseWriter, r *http.Request) {
	handler.GetAllWorkouts(a.DB, w, r)
}

func (a *App) CreateWorkout(w http.ResponseWriter, r *http.Request) {
	handler.CreateWorkout(a.DB, w, r)
}

func (a *App) GetWorkout(w http.ResponseWriter, r *http.Request) {
	handler.GetWorkout(a.DB, w, r)
}

func (a *App) UpdateWorkout(w http.ResponseWriter, r *http.Request) {
	handler.UpdateWorkout(a.DB, w, r)
}

func (a *App) DeleteWorkout(w http.ResponseWriter, r *http.Request) {
	handler.DeleteWorkout(a.DB, w, r)
}

func (a *App) DisableWorkout(w http.ResponseWriter, r *http.Request) {
	handler.DisableWorkout(a.DB, w, r)
}

func (a *App) EnableWorkout(w http.ResponseWriter, r *http.Request) {
	handler.EnableWorkout(a.DB, w, r)
}

// Handlers to manage Session Data
func (a *App) GetAllSessions(w http.ResponseWriter, r *http.Request) {
	handler.GetAllSessions(a.DB, w, r)
}

func (a *App) CreateSession(w http.ResponseWriter, r *http.Request) {
	handler.CreateSession(a.DB, w, r)
}

func (a *App) GetSession(w http.ResponseWriter, r *http.Request) {
	handler.GetSession(a.DB, w, r)
}

func (a *App) UpdateSession(w http.ResponseWriter, r *http.Request) {
	handler.UpdateSession(a.DB, w, r)
}

func (a *App) DeleteSession(w http.ResponseWriter, r *http.Request) {
	handler.DeleteSession(a.DB, w, r)
}

// Run the app on it's router
func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}
