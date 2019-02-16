package models

type AuthResponse struct {
	AccessToken string
	Scope       string
	TokenType   string
}
