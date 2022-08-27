package src

import (
	"io/ioutil"
	"log"
	"txp/web-service-gin/src/data"
)

// App struct
type App struct {
	router *Router
}

func (a *App) discardLog() {
	log.SetOutput(ioutil.Discard)
}

// Init app
func (a *App) Init() {
	// discard log in production
	// a.discardLog()
	data.Init()
	a.router = &Router{}
	a.router.Init()
}

// Run app
func (a *App) Run() {
	a.router.run()
}
