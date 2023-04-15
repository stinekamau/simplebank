package token

import (
	"time"

	"aidanwoods.dev/go-paseto"
)

type PasetoMaker struct {
	token        paseto.Token
	symmetricKey paseto.V3SymmetricKey
}

func NewPasetoMaker() (Maker, error) {

	return PasetoMaker{
		paseto.NewToken(),
		paseto.NewV3SymmetricKey(),
	}, nil

}

func (maker PasetoMaker) VerifyToken(token string) (*Payload, error) {

	// obtain the token and decrypt

	var payload *Payload

	parser := paseto.NewParserForValidNow()

	parser.AddRule(paseto.ForAudience("bank_clients"))
	parser.AddRule(paseto.IssuedBy("stanbic"))

	tkn, err := parser.ParseV3Local(maker.symmetricKey, token, nil)
	if err != nil {
		return nil, err
	}

	if err := tkn.Get("payload", &payload); err != nil {
		return nil, err
	}
	return payload, nil
}

func (maker PasetoMaker) CreateToken(username string, duration time.Duration) (string, error) {
	payload, err := NewPayload(username, duration)
	if err != nil {
		return "", err
	}

	// Add the payload to the token
	if err := maker.token.Set("payload", payload); err != nil {
		return "", err
	}

	// Set claims
	maker.token.SetAudience("bank_clients")
	maker.token.SetIssuedAt(time.Now())
	maker.token.SetExpiration(time.Now().Add(50 * time.Second))
	maker.token.SetNotBefore(time.Now())
	maker.token.SetIssuer("stanbic")

	encrypted := maker.token.V3Encrypt(maker.symmetricKey, nil)

	return encrypted, nil

}
