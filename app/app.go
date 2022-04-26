package app

import (
	"database/sql"
	"hades/app/handler"
	"hades/util"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// App has router and db instances
type App struct {
	Router *mux.Router
	DB     *util.SQLiteRepository
}

const fileName = "sqlite.db"

// Initialize initializes the app with predefined configuration
func (a *App) Initialize() {
	db, err := sql.Open("sqlite3", fileName)
	if err != nil {
		log.Fatal(err)
	}

	websiteRepository := util.NewSQLiteRepository(db)

	if err := websiteRepository.Migrate(); err != nil {
		log.Fatal(err)
	}
	a.DB = websiteRepository
	a.Router = mux.NewRouter()
	a.setRouters()
}

// setRouters sets the all required routers
func (a *App) setRouters() {
	// Routing for handling the projects
	a.Get("/websites", a.handleRequest(handler.GetAllWebsites))
	a.Get("/websites/{id}", a.handleRequest(handler.GetWebsite))
	a.Get("/websites/{website_id}/statuses", a.handleRequest(handler.GetWebsiteStatuses))

	//a.Post("/websites", a.handleRequest(handler.CreateProject))
	//a.Put("/projects/{title}", a.handleRequest(handler.UpdateProject))
	//a.Delete("/projects/{title}", a.handleRequest(handler.DeleteProject))
	//a.Delete("/projects/{title}/archive", a.handleRequest(handler.RestoreProject))

	// Routing for handling the tasks
	//a.Get("/projects/{title}/tasks", a.handleRequest(handler.GetAllTasks))
	//a.Post("/projects/{title}/tasks", a.handleRequest(handler.CreateTask))
	//a.Get("/projects/{title}/tasks/{id:[0-9]+}", a.handleRequest(handler.GetTask))
	//a.Put("/projects/{title}/tasks/{id:[0-9]+}", a.handleRequest(handler.UpdateTask))
	//a.Delete("/projects/{title}/tasks/{id:[0-9]+}", a.handleRequest(handler.DeleteTask))
	//a.Put("/projects/{title}/tasks/{id:[0-9]+}/complete", a.handleRequest(handler.CompleteTask))
	//a.Delete("/projects/{title}/tasks/{id:[0-9]+}/complete", a.handleRequest(handler.UndoTask))
}

// Get wraps the router for GET method
func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}

// Post wraps the router for POST method
func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}

// Put wraps the router for PUT method
func (a *App) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("PUT")
}

// Delete wraps the router for DELETE method
func (a *App) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("DELETE")
}

// Run the app on it's router
func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}

type RequestHandlerFunction func(db *util.SQLiteRepository, w http.ResponseWriter, r *http.Request)

func (a *App) handleRequest(handler RequestHandlerFunction) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(a.DB, w, r)
	}
}
