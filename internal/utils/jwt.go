package utils

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte(os.Getenv("JWT_secret"))

type JWTClaim struct {
	UserID uint   `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

type TokenPayload struct {
	UserID uint
	Role   string
}

func GenerateToken(payload TokenPayload) (string, error) {
	expirationTime := time.Now().Add(2 * time.Hour)

	claims := &JWTClaim{
		UserID: payload.UserID,
		Role:   payload.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(jwtSecret)
}

func ValidateToken(tokenString string) (*JWTClaim, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaim{}, func(t *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*JWTClaim)
	if !ok || !token.Valid {
		return nil, errors.New("Invalid token")
	}

	return claims, nil

}
