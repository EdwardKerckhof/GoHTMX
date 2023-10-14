package paseto

import (
	"fmt"
	"time"

	"github.com/aead/chacha20poly1305"
	"github.com/o1egl/paseto"

	"github.com/EdwardKerckhof/gohtmx/pkg/token"
)

type PasetoMaker struct {
	paseto       *paseto.V2
	symmetricKey []byte
}

func NewMaker(symmetricKey string) (token.Maker, error) {
	if len(symmetricKey) != chacha20poly1305.KeySize {
		return nil, fmt.Errorf("invalid key size: must be %d characters", chacha20poly1305.KeySize)
	}

	maker := &PasetoMaker{
		paseto:       paseto.NewV2(),
		symmetricKey: []byte(symmetricKey),
	}
	return maker, nil
}

func (m *PasetoMaker) GenerateToken(username string, duration time.Duration) (string, error) {
	payload := token.NewPayload(username, duration)
	return m.paseto.Encrypt(m.symmetricKey, payload, nil)
}

func (m *PasetoMaker) VerifyToken(t string) (*token.Payload, error) {
	payload := &token.Payload{}
	err := m.paseto.Decrypt(t, m.symmetricKey, payload, nil)
	if err != nil {
		return nil, token.ErrInvalidToken
	}

	err = payload.Valid()
	if err != nil {
		return nil, err
	}

	return payload, nil
}
