package routing

import (
	"fmt"
	"net/http"
	"strings"

	echo "github.com/labstack/echo/v4"
	"github.com/ravipwaghmare/MatchMove/interfaces"
	"github.com/ravipwaghmare/MatchMove/model"
)

// User struct
type User struct {
	userController  interfaces.User
	tokenController interfaces.Token
}

// NewUserRouter function
func NewUserRouter(
	userController interfaces.User,
	tokenController interfaces.Token,
) User {
	return User{
		userController:  userController,
		tokenController: tokenController,
	}
}

// Register new signin router
func (router User) Register(group *echo.Group) {
	group.POST("/signin", router.signIn)
	group.POST("/signup", router.signUp)
}

func (router User) signIn(context echo.Context) error {
	type request struct {
		Email    string `json:"email" validator:"required,email"`
		Password string `json:"password"`
		Token    string `json:"token"`
	}

	var req request

	// bind context to request
	if err := context.Bind(&req); err != nil {
		return err
	}

	// find user in database by email
	user, err := router.userController.FindByEmail(strings.TrimSpace(req.Email))
	if err != nil {
		return err
	}

	if user.AccType != "Admin" {
		if req.Token == "" {
			return fmt.Errorf("user can not access match move app unless user gets access token from admin")
		}
		if router.tokenController.ValidateToken(user, req.Token) {
			return fmt.Errorf("unable to login, Incorrect or invalid token")
		}
	}

	return context.JSON(http.StatusOK, model.SignInResponse{
		User: *user,
	})

}

func (router User) signUp(context echo.Context) error {
	type request struct {
		UserName string `json:"username" validator:"required"`
		Name     string `json:"name" validator:"required"`
		Email    string `json:"email" validator:"required,email"`
		Password string `json:"password" validator:"required"`
		AccType  string `json:"acctype" validator:"required"`
	}

	var req request

	// bind context to request
	if err := context.Bind(&req); err != nil {
		return err
	}

	user := &model.User{
		UserName: req.UserName,
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
		AccType:  req.AccType,
	}

	// Add user in the database
	err := router.userController.Create(user)
	if err != nil {
		return err
	}

	return context.JSON(http.StatusOK, map[string]string{
		"message": "User Added Successfully",
	})

}
