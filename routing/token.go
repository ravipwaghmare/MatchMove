package routing

import (
	"fmt"
	"net/http"

	"github.com/ravipwaghmare/MatchMove/interfaces"

	echo "github.com/labstack/echo/v4"
)

// Token struct
type Token struct {
	tokenController interfaces.Token
	userController  interfaces.User
}

// NewTokenRouter function
func NewTokenRouter(
	tokenController interfaces.Token,
	userController interfaces.User,

) Token {
	return Token{
		tokenController: tokenController,
		userController:  userController,
	}
}

// Register new token router
func (router Token) Register(group *echo.Group) {
	group.POST("/sendtoken", router.sendToken)
}

func (router Token) sendToken(context echo.Context) error {
	type request struct {
		ClientEmail string `json:"client_email" validator:"required"`
		AdminEmail  string `json:"admin_email" validator:"required"`
	}

	var req request

	// bind context to request
	if err := context.Bind(&req); err != nil {
		return err
	}

	adminUsr, err := router.userController.FindByEmail(req.AdminEmail)
	if err != nil {
		return err
	}

	if adminUsr.AccType != "Admin" {
		return fmt.Errorf("only admins are allowed to send token")
	}

	clientUsr, err := router.userController.FindByEmail(req.ClientEmail)
	if err != nil {
		return err
	}

	if clientUsr.AccType == "Admin" {
		return fmt.Errorf("sending token to admin user is not allowed")
	}

	// Add weight in the database
	err = router.tokenController.SendTokenToClient(adminUsr, clientUsr)
	if err != nil {
		return err
	}

	return context.JSON(http.StatusOK, map[string]string{
		"message": "Token sent successfully",
	})
}
