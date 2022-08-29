package src

import (
	"txp/web-service-gin/src/core/middleware"
	"txp/web-service-gin/src/module/content"
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
	r.registerContentRoutes(
		util.V1,
	)
}

func (r *Router) registerMiddlewares() {
	r.Engine.Use(
		middleware.JSONMiddleware,
		middleware.ErrorMiddleware,
	)
}

func (r *Router) registerUserRoutes(
	version string,
) {
	h := &user.UserHandler{}
	h.InitDependencies()
	group := r.Engine.Group(
		util.ApiPattern + version + util.UsersPattern,
	)
	{
		group.GET(
			util.RootPattern,
			h.FindAll,
		)
		group.GET(
			util.RootPattern+":id",
			h.FindOne,
		)
		group.POST(
			util.RootPattern,
			h.Create,
		)
		group.PATCH(
			util.RootPattern+":id",
			h.Update,
		)
		group.DELETE(
			util.RootPattern+":id",
			h.Delete,
		)
	}
}

func (r *Router) registerContentRoutes(
	version string,
) {
	h := &content.ContentHandler{}
	h.InitDependencies()
	group := r.Engine.Group(
		util.ApiPattern + version + util.ContentsPattern,
	)
	{
		group.GET(
			util.RootPattern,
			h.FindAll,
		)
		group.GET(
			util.RootPattern+":id",
			h.FindOne,
		)
		group.POST(
			util.RootPattern,
			h.Create,
		)
		group.PATCH(
			util.RootPattern+":id",
			h.Update,
		)
		group.DELETE(
			util.RootPattern+":id",
			h.Delete,
		)
	}
}

func (r *Router) run() {
	r.Engine.Run(":8080")
}
