package token

import "time"

// Maker interface for managing token
type Maker interface {
	// CreateToken creates a now token for specific username and duration
	CreateToken(username string, duration time.Duration) (string, *Payload, error)

	// VerifyToken checks if the token is valid or not 
	VerifyToken(token string) (*Payload, error)
}