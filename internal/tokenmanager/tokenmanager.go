package tokenmanager

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"math/rand"
	"time"
)

type TokenManager interface {
	NewJWT(userName string) (string, error)
	Parse(accessToken string) (string, error)
	NewRefreshToken() (string, error)
}

type Manager struct {
	SignKey string
}

func New(signingKey string) (TokenManager, error) {
	if signingKey == "" {
		return nil, fmt.Errorf("empty signing key")
	}
	return &Manager{SignKey: signingKey}, nil
}

func (t *Manager) NewJWT(userName string) (string, error) {
	// create new JWT Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		Subject:   userName,
	})
	// Sign and get the complete encoded token as a string using the secret
	return token.SignedString([]byte(t.SignKey))

}

func (t *Manager) Parse(accessToken string) (string, error) {
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(t.SignKey), nil
	})
	if err != nil {
		return "", err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", fmt.Errorf("error get user claims from token")
	}

	return claims["sub"].(string), nil
}

func (t *Manager) NewRefreshToken() (string, error) {
	b := make([]byte, 32)

	src := rand.NewSource(time.Now().Unix())
	r := rand.New(src)

	_, err := r.Read(b)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
