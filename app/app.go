package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
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
	// a.DB = model.DBMigrateWorkoutFacility(db)
	a.DB = model.DBMigrateFacility(db)
	a.DB = model.DBMigrateLists(db)
	a.Router = mux.NewRouter()
	a.setRouters()
}

// Set all required routers
func (a *App) setRouters() {
	// Workout
	a.Get("/workouts/all", a.GetAllWorkouts)
	a.Get("/workouts/public", a.GetPublicWorkouts)
	a.Get("/workouts", a.GetUserWorkouts) // Workouts per User
	a.Post("/workouts", a.CreateWorkout)
	a.Get("/workouts/{id}", a.GetWorkout)
	a.Put("/workouts/{id}", a.UpdateWorkout)
	a.Delete("/workouts/{id}", a.DeleteWorkout)
	a.Put("/workouts/{id}/disable", a.DisableWorkout)
	a.Put("/workouts/{id}/enable", a.EnableWorkout)

	// Session
	a.Get("/sessions", a.GetAllUserSessions)                 // Sessions per User
	a.Get("/sessions/details", a.GetUserSessionsWithDetails) // Sessions per User
	a.Post("/sessions", a.CreateSession)
	a.Get("/sessions/{id}", a.GetSession)
	a.Put("/sessions/{id}", a.UpdateSession)
	a.Delete("/sessions/{id}", a.DeleteSession)

	// WorkoutExecution
	a.Get("/executions", a.GetAllExecutions) // Executions per Session (and per User)
	a.Post("/executions", a.CreateExecution)
	a.Get("/executions/{id}", a.GetExecution)
	// a.Put("/executions/{id}", a.UpdateExecution)
	// a.Delete("/executions/{id}", a.DeleteExecution)

	// Lists
	a.Get("/list/facilities", a.GetAllFacilities)
	a.Get("/list/energysystems", a.GetAllEnergySystems)
	a.Get("/list/experiencelevels", a.GetAllExperienceLevels)
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

func (a *App) GetPublicWorkouts(w http.ResponseWriter, r *http.Request) {
	handler.GetPublicWorkouts(a.DB, w, r)
}

func (a *App) GetUserWorkouts(w http.ResponseWriter, r *http.Request) {
	handler.GetUserWorkouts(a.DB, w, r)
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
func (a *App) GetAllUserSessions(w http.ResponseWriter, r *http.Request) {
	handler.GetAllUserSessions(a.DB, w, r)
}

func (a *App) GetUserSessionsWithDetails(w http.ResponseWriter, r *http.Request) {
	handler.GetUserSessionsWithDetails(a.DB, w, r)
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

// Handlers to manage Execution Data
func (a *App) GetAllExecutions(w http.ResponseWriter, r *http.Request) {
	handler.GetAllExecutions(a.DB, w, r)
}

func (a *App) CreateExecution(w http.ResponseWriter, r *http.Request) {
	handler.CreateExecution(a.DB, w, r)
}

func (a *App) GetExecution(w http.ResponseWriter, r *http.Request) {
	handler.GetExecution(a.DB, w, r)
}

// Handlers to manage List Data
func (a *App) GetAllFacilities(w http.ResponseWriter, r *http.Request) {
	handler.GetAllFacilities(a.DB, w, r)
}

func (a *App) GetAllEnergySystems(w http.ResponseWriter, r *http.Request) {
	handler.GetAllEnergySystems(a.DB, w, r)
}

func (a *App) GetAllExperienceLevels(w http.ResponseWriter, r *http.Request) {
	handler.GetAllExperienceLevels(a.DB, w, r)
}

// Run the app on its router
func (a *App) Run(host string) {
	// https://www.thepolyglotdeveloper.com/2017/10/handling-cors-golang-web-application/
	log.Fatal(http.ListenAndServe(host, handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(a.Router)))
	// log.Fatal(http.ListenAndServe(":3000", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(router)))
}
