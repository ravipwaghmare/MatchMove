package controllers

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"net/smtp"
	"time"

	"github.com/ravipwaghmare/MatchMove/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// Token controller
type Token struct {
	database *gorm.DB
}

// NewTokenController will create new Token controller
func NewTokenController(database *gorm.DB) Token {
	return Token{
		database: database,
	}
}

func (controller Token) ValidateToken(user *model.User, token string) bool {

	tokenMap := new(model.TokenMapping)

	if err := controller.database.Where("client_id = ? AND token = ?", user.ID, token).
		First(tokenMap).Order(clause.OrderByColumn{Column: clause.Column{Name: "id"}, Desc: true}).Error; err != nil {
		return true
	}

	expiredAtDate, err := time.Parse("02-Jan-2006 15:04:05", tokenMap.ExpiredAt)
	if err != nil {
		return false
	}

	currentDate, err := time.Parse("02-Jan-2006 15:04:05", time.Now().Format("02-Jan-2006 15:04:05"))
	if err != nil {
		return false
	}

	if !expiredAtDate.Before(currentDate) {
		return false
	}

	return true
}

func (controller Token) SendTokenToClient(adminUsr, clientUsr *model.User) error {

	tokenMap := new(model.TokenMapping)

	var token string

	if err := controller.database.Where("admin_id = ? And client_id = ?", adminUsr.ID, clientUsr.ID).
		First(tokenMap).Order(clause.OrderByColumn{Column: clause.Column{Name: "id"}, Desc: true}).Error; err != nil {
		token = generateSecureToken(8)
		fmt.Println("Entered here")
		tokenMap.Token = token
		tokenMap.AdminID = adminUsr.ID
		tokenMap.ClientID = clientUsr.ID

		expiredAt := time.Now().Add(48 * time.Hour).Format("02-Jan-2006 15:04:05")

		tokenMap.ExpiredAt = expiredAt
		err := controller.database.Create(tokenMap).Error
		if err != nil {
			return err
		}
	} else {

		expiredAtDate, err := time.Parse("02-Jan-2006 15:04:05", tokenMap.ExpiredAt)
		if err != nil {
			return err
		}

		currentDate, err := time.Parse("02-Jan-2006 15:04:05", time.Now().Format("02-Jan-2006 15:04:05"))
		if err != nil {
			return err
		}

		if expiredAtDate.Before(currentDate) {
			tokenMap.Token = token
			tokenMap.AdminID = adminUsr.ID
			tokenMap.ClientID = clientUsr.ID

			expiredAt := time.Now().Add(48 * time.Hour).Format("02-Jan-2006 15:04:05")

			tokenMap.ExpiredAt = expiredAt
			err := controller.database.Create(tokenMap).Error
			if err != nil {
				return err
			}
		}
	}

	token = tokenMap.Token

	// Sender data.
	from := adminUsr.Email
	password := "admin_password"

	// Receiver email address.
	to := []string{
		clientUsr.Email,
	}

	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Message.
	message := []byte("Access MatchMove App using following token :" + token)

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func generateSecureToken(length int) string {
	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return hex.EncodeToString(b)
}
