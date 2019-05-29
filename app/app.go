package app
 
import (
	"fmt"
	"log"
	"net/http"
 
	"github.com/weyermann/trainr-api/app/handler"
	"github.com/weyermann/trainr-api/app/model"
	"github.com/weyermann/trainr-api/config"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
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
	a.DB = model.DBMigrateWorkoutExecutionInfo(db)
	a.Router = mux.NewRouter()
	a.setRouters()
}
 
// Set all required routers
func (a *App) setRouters() {
	// Routing for handling the projects
	a.Get("/workouts", a.GetAllWorkouts)
	a.Post("/workouts", a.CreateWorkout)
	a.Get("/workouts/{title}", a.GetWorkout)
	a.Put("/workouts/{title}", a.UpdateWorkout)
	a.Delete("/workouts/{title}", a.DeleteWorkout)
	a.Put("/workouts/{title}/disable", a.DisableWorkout)
	a.Put("/workouts/{title}/enable", a.EnableWorkout)
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
 
// Run the app on it's router
func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}