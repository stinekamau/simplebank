package token

import "time"

type Maker interface {
	VerifyToken(token string) (*Payload, error)
	CreateToken(username string, duration time.Duration) (string, error)
}
