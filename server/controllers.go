package server

import (
	"github.com/ravipwaghmare/MatchMove/controllers"

	"gorm.io/gorm"
)

var (
	UserController  controllers.User
	TokenController controllers.Token
)

//InitializeControllers initialize the controllers to communicate
func InitializeControllers(db *gorm.DB) {
	UserController = controllers.NewUserController(db)
	TokenController = controllers.NewTokenController(db)
}
