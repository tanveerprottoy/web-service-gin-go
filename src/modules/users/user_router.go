package users

import "Router"

func (router *Router) registerUserRoutes(router *Router) {
		userHandler := &users.UserHandler{}
	router.Engine.Group(
		BasePattern + UserPattern,
		func(r chi.Router) {
			r.Post("/login", authHandler.Login)
			r.Post("/pass-auth", authHandler.PassAuth)
		},
	)
}