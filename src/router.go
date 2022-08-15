package app

import (
	"net/http"

	"biyelap.com/biyelap-core/app/handler"
	"github.com/go-chi/chi"
)

// Router struct
type Router struct {
	Mux  *chi.Mux
	TMux *http.ServeMux
}

// BasePattern path
const BasePattern = "/v1/api"
const RootPattern = "/"
const AuthPattern = "/auth"
const FavouritesPattern = "/favourites"
const PreferencesPattern = "/preferences"
const ConnectionsPattern = "/connections"
const ConnectionRequestsPattern = "/connection-requests"
const IdPathPattern = "/{id}"
const FavIdPathPattern = "/{fav_id}"
const ConnIdPathPattern = "/{conn_id}"
const ConnReqIdPathPattern = "/{conn_req_id}"
const PreferenceIdPathPattern = "/{preference_id}"

func (router *Router) Init() {
	router.Mux = chi.NewRouter()
	router.registerRoutes()
}

func (router *Router) registerRoutes() {
	authHandler := &handler.AuthHandler{}
	router.registerGlobalMiddleWares()
	router.registerPublicRoutes()
	router.registerAuthRoutes(authHandler)
	router.registerProtectedRoutes(authHandler)
}

func (router *Router) registerAuthRoutes(authHandler *handler.AuthHandler) {
	// auth routes non protected
	router.Mux.Route(
		BasePattern+AuthPattern,
		func(r chi.Router) {
			r.Post("/login", authHandler.Login)
			r.Post("/pass-auth", authHandler.PassAuth)
		},
	)
}

func (router *Router) registerProtectedRoutes(authHandler *handler.AuthHandler) {
	// protected routes
	router.Mux.Group(
		func(r chi.Router) {
			fileHandler := &handler.FileHandler{}
			r.Use(JWTMiddleWare)
			r.Post(BasePattern+AuthPattern+RootPattern, authHandler.RefreshToken)
			r.Post(BasePattern+"/files", fileHandler.PostFile)
			// user routes and sub routes
			r.Route(
				BasePattern+"/users",
				func(r chi.Router) {
					userHandler := &handler.UserHandler{}
					r.Get(RootPattern, userHandler.GetUsers)
					// Sub routers
					r.Route(
						IdPathPattern,
						func(r chi.Router) {
							favouriteHandler := &handler.FavouriteHandler{}
							preferenceHandler := &handler.PreferenceHandler{}
							connectionHandler := &handler.ConnectionHandler{}
							connectionRequestHandler := &handler.ConnectionRequestHandler{}
							r.Get(RootPattern, userHandler.GetUser)
							r.Put(RootPattern, userHandler.PutUser)
							r.Get(FavouritesPattern, favouriteHandler.GetFavourites)
							r.Post(FavouritesPattern, favouriteHandler.PostFavourite)
							r.Get(PreferencesPattern, preferenceHandler.GetPreference)
							r.Post(PreferencesPattern, preferenceHandler.PostPreference)
							r.Get(ConnectionsPattern, connectionHandler.GetConnections)
							// r.Post(ConnectionsPattern, connectionHandler.PostConnection)
							r.Get(ConnectionRequestsPattern, connectionRequestHandler.GetConnectionRequests)
							r.Post(ConnectionRequestsPattern, connectionRequestHandler.PostConnectionRequest)
							// fav sub routers
							r.Route(
								FavouritesPattern+FavIdPathPattern,
								func(r chi.Router) {
									r.Delete(RootPattern, favouriteHandler.DeleteFavourite)
								},
							)
							// interest sub routers
							r.Route(
								PreferencesPattern+PreferenceIdPathPattern,
								func(r chi.Router) {
									r.Put(RootPattern, preferenceHandler.PutPreference)
								},
							)
							// connection sub routers
							r.Route(
								ConnectionsPattern+ConnIdPathPattern,
								func(r chi.Router) {
									r.Delete(RootPattern, connectionHandler.DeleteConnection)
								},
							)
							// connection requests sub routers
							r.Route(
								ConnectionRequestsPattern+ConnReqIdPathPattern,
								func(r chi.Router) {
									r.Put(RootPattern, connectionRequestHandler.PutConnectionRequest)
									r.Delete(RootPattern, connectionRequestHandler.DeleteConnectionRequest)
								},
							)
						},
					)
				},
			)
		},
	)
}

func (router *Router) registerPublicRoutes() {
	// package routes
	/*router.Mux.Route(
		"/packages",
		func(r chi.Router) {
			packageHandler := &handler.PackageHandler{}
			r.Get("/", packageHandler.GetPackages)
		},
	)*/
}
