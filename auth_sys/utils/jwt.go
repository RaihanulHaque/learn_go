package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("your_secret_key") // Replace with an env var in production

func GenerateJWT(username string) (string, error) {
    claims := &jwt.RegisteredClaims{
        Subject:   username,
        ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(jwtKey)
}

func ValidateJWT(tokenStr string) (string, error) {
    token, err := jwt.ParseWithClaims(tokenStr, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
        return jwtKey, nil
    })

    if err != nil {
        return "", err
    }

    if claims, ok := token.Claims.(*jwt.RegisteredClaims); ok && token.Valid {
        return claims.Subject, nil
    }

    return "", jwt.ErrTokenSignatureInvalid
}
