package interfaces

import "github.com/ravipwaghmare/MatchMove/model"

// Token interface
type Token interface {
	SendTokenToClient(adminUsr, clientUsr *model.User) error
	ValidateToken(user *model.User, token string) bool
}
