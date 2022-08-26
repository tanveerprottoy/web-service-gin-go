package app

import (
	"txp/web-service-gin/src/core/middleware"

	"github.com/gin-gonic/gin"
)

// Router struct
type Router struct {
	Engine gin.Engine
}

func (router *Router) Init() {
	router.Engine = *gin.Default()
	router.registerMiddlewares()
	router.registerRoutes()
}

func (router *Router) registerMiddlewares() {
	router.Engine.Use(
		middleware.ErrorHandler,
	)
}

func (router *Router) registerRoutes() {

}
