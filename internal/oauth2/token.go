package oauth2

import "time"

type Token struct {
	AccessToken  string
	RefreshToken string
	TokenType    string
	ClientID     string
	UserID       string
	Scopes       []string
	Expiry       time.Time
}
