package model

type TokenMapping struct {
	ID        uint   `json:"id" gorm:"primary_key"`
	AdminID   uint   `json:"admin_id"`
	ClientID  uint   `json:"client_id"`
	Token     string `json:"token"`
	ExpiredAt string `json:"expired_at"`
}
