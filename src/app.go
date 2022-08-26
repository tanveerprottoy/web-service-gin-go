package src

import (
	"io/ioutil"
	"log"
)

// App struct
type App struct {
	router *Router
}

func (app *App) discardLog() {
	log.SetOutput(ioutil.Discard)
}

// Init app
func (app *App) Init() {
	// discard log in production
	app.discardLog()
	// data.Init()
	app.router = &Router{}
	app.router.Init()
}

// Run app
func (app *App) Run() {
	app.router.run()
}
