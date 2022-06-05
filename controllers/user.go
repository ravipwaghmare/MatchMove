package controllers

import (
	"strings"

	"github.com/ravipwaghmare/MatchMove/model"

	"gorm.io/gorm"
)

// User controller
type User struct {
	database *gorm.DB
}

// NewUserController will create new user controller
func NewUserController(database *gorm.DB) User {
	return User{
		database: database,
	}
}

// Create will make a new user
func (controller User) Create(user *model.User) error {
	user.Email = strings.ToLower(strings.TrimSpace(user.Email))
	return controller.database.Create(user).Error
}

// FindByEmail will find user by email
func (controller User) FindByEmail(email string) (*model.User, error) {
	user := new(model.User)
	if err := controller.database.Where("email = ?", strings.ToLower(email)).First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
