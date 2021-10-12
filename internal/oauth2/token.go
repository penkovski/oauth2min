package oauth2

type Token struct {
	AccessToken  string   `json:"access_token"`
	RefreshToken string   `json:"refresh_token"`
	TokenType    string   `json:"token_type"`
	ClientID     string   `json:"client_id"`
	UserID       string   `json:"user_id"`
	Scopes       []string `json:"scopes"`
	Expiry       int64    `json:"expiry"`
}
