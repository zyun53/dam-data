package app

import (
	"log"
	"net/http"

	"github.com/zyun-i/dam-data/config"
	"github.com/zyun-i/dam-data/handler"
	"github.com/zyun-i/dam-data/model"

	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// App has router and db instances
type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

// App initialize with predefined configuration
func (a *App) Initialize(config *config.Config) {
	dsn := "host=localhost user=postgres password=pass dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Tokyo"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Could not connect database")
	}

	a.DB = model.DBMigrate(db)
	a.Router = mux.NewRouter()
	a.setRouters()
}

// Set all required routers
func (a *App) setRouters() {
	// Routing for handling the projects
	a.Get("/dams", a.GetAllDams)
	a.Get("/dams/{id:[0-9]+}", a.GetDam)
	a.Get("/nldams", a.GetAllNlDams)
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

// Handlers to manage Employee Data
func (a *App) GetAllDams(w http.ResponseWriter, r *http.Request) {
	handler.GetAllDams(a.DB, w, r)
}

func (a *App) GetDam(w http.ResponseWriter, r *http.Request) {
	handler.GetDam(a.DB, w, r)
}

// Run the app on it's router
func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}

// Handlers to manage Employee Data
func (a *App) GetAllNlDams(w http.ResponseWriter, r *http.Request) {
	handler.GetAllNlDams(a.DB, w, r)
}
