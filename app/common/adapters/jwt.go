package adapters

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type jwtAdapter struct {
	key           []byte
	tokenDuration time.Duration
}

type claims struct {
	Sum string
	jwt.StandardClaims
}

func NewJWTAdapter(key string, dur time.Duration) jwtAdapter {
	return jwtAdapter{[]byte(key), dur}
}

func (j jwtAdapter) Encrypt(id string) (string, error) {
	now := time.Now()

	claim := &claims{Sum: id, StandardClaims: jwt.StandardClaims{
		ExpiresAt: now.Add(j.tokenDuration).Unix(),
		IssuedAt:  now.Unix(),
	}}

	tk := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	signedToken, err := tk.SignedString(j.key)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (j jwtAdapter) Decrypt(val string) (string, error) {
	claims := &claims{}
	token, err := jwt.ParseWithClaims(val, claims, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid token: %v", val)
		}
		return j.key, nil
	})
	if err != nil || !token.Valid {
		return "", fmt.Errorf("unauhtorized: %v", token)
	}

	return claims.Sum, nil
}
