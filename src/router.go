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

func (r *Router) Init() {
	r.Engine = gin.Default()
	r.registerMiddlewares()
	r.registerUserRoutes(
		util.V1,
	)
}

func (r *Router) registerMiddlewares() {
	r.Engine.Use(
		middleware.ErrorHandler,
	)
}

func (r *Router) registerUserRoutes(
	version string,
) {
	handler := &user.UserHandler{}
	handler.InitDependencies()
	group := r.Engine.Group(
		util.ApiPattern + version + util.UsersPattern,
	)
	{
		group.GET(
			util.RootPattern,
			handler.FindUsers,
		)
		group.GET(
			util.RootPattern+":id",
			handler.FindUser,
		)
		group.POST(
			util.RootPattern,
			handler.CreateUser,
		)
		group.PATCH(
			util.RootPattern+":id",
			handler.UpdateUser,
		)
		group.DELETE(
			util.RootPattern+":id",
			handler.DeleteUser,
		)
	}
}

func (r *Router) run() {
	r.Engine.Run(":8080")
}
