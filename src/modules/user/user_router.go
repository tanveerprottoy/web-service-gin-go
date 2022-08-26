package user

import (
	app "txp/web-service-gin/src"
	"txp/web-service-gin/src/util"
)

func registerUserRoutes(
	router *app.Router,
	version string,
) {
	handler := &UserHandler{}
	group := router.Engine.Group(
		util.ApiPattern + version + util.UserPattern,
	)
	{
		group.GET(
			util.RootPattern,
			handler.GetUsers,
		)
		group.GET(
			util.RootPattern + ":id",
			handler.GetUser,
		)
	}
}
