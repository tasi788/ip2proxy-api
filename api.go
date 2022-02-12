package main

import (
	"github.com/glebarez/sqlite"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"net/http"
)

type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

func (a *App) Initialize() {
	db, _ := gorm.Open(sqlite.Open("sqlite.db"), &gorm.Config{})

	a.DB = DBMigrate(db)
	a.Router = mux.NewRouter()
	a.setRouters()

}



// Set all required routers
func (a *App) setRouters() {
	// Routing for handling the projects
	a.Get("/add", a.addUser)
	a.Get("/query", a.query)
	a.Delete("/delete", a.deleteUser)
	//a.Get("/employees/{title}", a.GetEmployee)
}

// Get Wrap the router for GET method
func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}

// Post Wrap the router for POST method
func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}

// Delete Wrap the router for POST method
func (a *App) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("DELETE")
}

// Router

func (a *App) addUser(w http.ResponseWriter, r *http.Request) {
	addUser(a.DB, w, r)
}

func (a *App) query(w http.ResponseWriter, r *http.Request) {
	query(a.DB, w, r)
}

func (a *App) deleteUser(w http.ResponseWriter, r *http.Request) {
	deleteUser(a.DB, w, r)
}