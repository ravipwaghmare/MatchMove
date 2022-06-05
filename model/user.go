package model

// User model in DB
type User struct {
	ID       uint   `json:"user_id" gorm:"primary_key"`
	UserName string `json:"username"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	AccType  string `json:"acctype"`
}

// SignInResponse response body
type SignInResponse struct {
	User User `json:"user"`
}
