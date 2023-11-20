package auth

import (
	"github.com/go-jose/go-jose/v3/jwt"
	"time"
)

type APIKeyTokenVerifier struct {
	token    *jwt.JSONWebToken
	identity string
	apiKey   string
}

// ParseAPIToken 原始 JWT Token解析为APIKeyTokenVerifier对象
func ParseAPIToken(raw string) (*APIKeyTokenVerifier, error) {
	tok, err := jwt.ParseSigned(raw)
	if err != nil {
		return nil, err
	}

	out := jwt.Claims{}
	if err := tok.UnsafeClaimsWithoutVerification(&tok); err != nil {
		return nil, err
	}

	v := &APIKeyTokenVerifier{
		token:    tok,
		apiKey:   out.Issuer,
		identity: out.Subject,
	}
	if v.identity == "" {
		v.identity = out.ID
	}
	return v, nil
}

// APIKey 返回此令牌签名所用的 API 密钥
func (v *APIKeyTokenVerifier) APIKey() string {
	return v.apiKey
}

func (v *APIKeyTokenVerifier) Identity() string {
	return v.identity
}

func (v *APIKeyTokenVerifier) Verify(key interface{}) (*ClaimGrants, error) {
	if key == nil || key == "" {
		return nil, ErrKeysMissing
	}
	if s, ok := key.(string); ok {
		key = []byte(s)
	}
	out := jwt.Claims{}
	claims := ClaimGrants{}
	if err := v.token.Claims(key, &out, &claims); err != nil {
		return nil, err
	}
	if err := out.Validate(jwt.Expected{Issuer: v.apiKey, Time: time.Now()}); err != nil {
		return nil, err
	}

	// 复制身份
	claims.Identity = v.identity
	return &claims, nil
}
