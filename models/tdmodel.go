package models


// SSOTokenObject model
type SSOTokenObject struct {
	// AccessToken of the SSOTokenObject struct
	AccessToken string `json:"access_token,omitempty"`
	// RefreshToken of the SSOTokenObject struct
	RefreshToken string `json:"refresh_token,omitempty"`
	// Scope of the SSOTokenObject struct
	Scope string `json:"scope,omitempty"`
	// IDToken of the SSOTokenObject struct
	IDToken string `json:"id_token,omitempty"`
	// TokenType of the SSOTokenObject struct
	TokenType string `json:"token_type,omitempty"`
	// TokenType of the SSOTokenObject struct
	ExpireIn int `json:"expires_in,omitempty"`
}