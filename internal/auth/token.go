package auth

import "time"

// Access tokens are signed with ES256 using keys from Secret Manager, rotated every 30 days.
// The public keys are served at /.well-known/jwks.json so other services verify tokens locally.
const (
	AccessTTL  = 15 * time.Minute
	RefreshTTL = 30 * 24 * time.Hour
)

// TokenRequest is the body of POST /v1/auth/token.
type TokenRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
