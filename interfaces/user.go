package interfaces

import (
	"github.com/ravipwaghmare/MatchMove/model"
)

// User interface
type User interface {
	Create(user *model.User) error
	FindByEmail(email string) (*model.User, error)
}
