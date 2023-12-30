package utils

import (
	"encoding/base64"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(ttl time.Duration, payload interface{}, secretKey string) (string, error) {
	decodeSecretKey, err := base64.StdEncoding.DecodeString(secretKey)
	if err != nil {
		return "", fmt.Errorf("could not decode secret key %w", err)
	}

	key, err := jwt.ParseRSAPrivateKeyFromPEM(decodeSecretKey)
	if err != nil {
		return "", fmt.Errorf("create: parse key: %w", err)
	}

	now := time.Now()
	claims := make(jwt.MapClaims)
	claims["sub"] = payload
	claims["exp"] = now.Add(ttl).Unix()
	claims["iat"] = now.Unix()
	claims["nbf"] = now.Unix()

	token, err := jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(key)
	if err != nil {
		return "", fmt.Errorf("create: sign token: %w", err)
	}

	return token, nil
}
