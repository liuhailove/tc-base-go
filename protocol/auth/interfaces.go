package auth

import "errors"

var (
	ErrKeysMissing = errors.New("missing API key or secret key")
)

type TokenVerifier interface {
	Identity() string
	Verify(key interface{}) (*ClaimGrants, error)
}
type KeyProvider interface {
	GetSecret(key string) string
	NumKeys() int
}
