package src

import (
	"txp/web-service-gin/src/core/middleware"
	"txp/web-service-gin/src/module/user"
	"txp/web-service-gin/src/util"

	"github.com/gin-gonic/gin"
)

// Router struct
type Router struct {
	Engine *gin.Engine
}

func (router *Router) Init() {
	router.Engine = gin.Default()
	router.registerMiddlewares()
	router.registerRoutes()
}

func (router *Router) registerMiddlewares() {
	router.Engine.Use(
		middleware.ErrorHandler,
	)
}

func (router *Router) registerRoutes() {
	registerUserRoutes(
		router,
		util.V1,
	)
}

func (router *Router) run() {
	router.Engine.Run(":8080")
}

func registerUserRoutes(
	router *Router,
	version string,
) {
	handler := &user.UserHandler{}
	handler.InitDependencies()
	group := router.Engine.Group(
		util.ApiPattern + version + util.UsersPattern,
	)
	{
		group.GET(
			util.RootPattern,
			handler.FindUsers,
		)
		/* group.GET(
			util.RootPattern + ":id",
			handler.FindUsers,
		) */
	}
}
