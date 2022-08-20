package app

import (
	"github.com/gin-gonic/gin"
)

// Router struct
type Router struct {
	Engine      gin.Engine
}

// BasePattern path
const BasePattern = "/v1/api"
const RootPattern = "/"
const UserPattern = "/users"

func (router *Router) Init() {
	router.Engine = *gin.Default()
	router.registerRoutes()
}

func (router *Router) registerRoutes() {
	
}


