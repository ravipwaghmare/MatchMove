package server

import (
	"github.com/ravipwaghmare/MatchMove/routing"
)

var (
	UserRouter  routing.User
	TokenRouter routing.Token
)

//InitializeRouters will create routing page for user
func InitializeRouters(server Server) {
	UserRouter = routing.NewUserRouter(
		UserController,
		TokenController,
	)

	TokenRouter = routing.NewTokenRouter(
		TokenController,
		UserController,
	)

	// register routes
	authGroup := server.Group("/auth")
	UserRouter.Register(authGroup.Group("/user"))

	TokenRouter.Register(server.Group("/token"))

}
