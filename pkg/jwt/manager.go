package jwt

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

type Manager struct {
	secret        string
	tokenDuration time.Duration
}

type UserClaims struct {
	jwt.StandardClaims
	UserUUID string `json:"user_uuid"`
}

func NewManager(secret string, tokenDuration time.Duration) *Manager {
	return &Manager{
		secret:        secret,
		tokenDuration: tokenDuration,
	}
}

func (m *Manager) Generate(userUUID string) (string, error) {
	claims := UserClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(m.tokenDuration).Unix(),
		},
		UserUUID: userUUID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(m.secret))
}

func (m *Manager) Verify(accessToken string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(accessToken, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("token signing method")
		}

		return []byte(m.secret), nil
	})

	if err != nil {
		return nil, fmt.Errorf("invalid token: %w", err)
	}

	claims, ok := token.Claims.(*UserClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}

	return claims, nil
}
