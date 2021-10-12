package server

import "io"

type TokenRequest struct {
	ClientID     string   `json:"client_id"`
	ClientSecret string   `json:"client_secret"`
	GrantType    string   `json:"grant_type"`
	RefreshToken string   `json:"refresh_token"`
	AuthCode     string   `json:"auth_code"`
	Username     string   `json:"username"`
	Password     string   `json:"password"`
	Scopes       []string `json:"scopes"`
	RedirectURI  string   `json:"redirect_uri"`
	TokenFormat  string   `json:"token_format"`
}

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token,omitempty"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int64  `json:"expires_in"`
	Scope        string `json:"scope,omitempty"`
	TokenFormat  string `json:"token_format,omitempty"`
}

func (r *TokenRequest) Decode(req io.Reader) error {
	return nil
}
