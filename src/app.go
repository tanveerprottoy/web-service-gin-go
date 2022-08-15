package app

import (
	"io/ioutil"
	"log"
	"github.com/gin-gonic/gin"
)

// App struct
type App struct {
	Router *Router
}

func (app *App) discardLog() {
	log.SetOutput(ioutil.Discard)
}

// Init app
func (app *App) Init() {
	// discard log in production
	app.discardLog()
	datum.Init()
	app.Router = &Router{}
	app.Router.Init()
	app.initCOS()
}

// Run app
func (app *App) Run() {
	log.Fatal(
		http.ListenAndServe(
			":3000",
			app.Router.Mux,
		),
	)
}
