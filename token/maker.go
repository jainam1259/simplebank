package token

import "time"

// Maker is an interface to manage tokens
type Maker interface {
	// CreateToken creates a new token for a specific username & duration
	CreateToken(username string, duration time.Duration) (string, error)

	// VerifyToken checks if the token is valid or not
	VerifyToken(token string) (*Payload, error)
}
