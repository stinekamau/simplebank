package token

import (
	"fmt"
	"time"
)

type JWTMaker struct {
	secretKey string
}

func (maker JWTMaker) CreateToken(username string, duration time.Duration) (string, error) {
	payload, err := NewPayload(username, duration)
	if err != nil {
		return "", err
	}
	fmt.Println(payload)

	// token := jwt.NewWithClaims(jwt.SigningMethodES384, payload)

	// return token.SignedString([]byte(payload.Username))
	return "", nil
}

func (maker JWTMaker) VerifyToken(token string) (*Payload, error) {

	return nil, nil
}

func NewJWTMaker(secret string) (Maker, error) {

	jMaker := JWTMaker{
		secretKey: secret,
	}

	return jMaker, nil

}
