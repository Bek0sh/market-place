package utils

import (
	"encoding/base64"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func CreateToken(duration int, payload interface{}, privateKey string) (string, error) {
	decodedKey, err := base64.StdEncoding.DecodeString(privateKey)
	if err != nil {
		return "", fmt.Errorf("failed to decode key, error: %s", err.Error())
	}

	key, err := jwt.ParseRSAPrivateKeyFromPEM(decodedKey)

	if err != nil {
		return "", fmt.Errorf("failed to parse decoded key, error: %s", err.Error())
	}

	claims := make(jwt.MapClaims)
	claims["sub"] = payload
	claims["exp"] = duration * int(time.Minute)

	signedToken, err := jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(key)

	if err != nil {
		return "", errors.New("failed to sign token")
	}

	return signedToken, nil
}

func VerifyToken(token, publicKey string) (interface{}, error) {
	decodedKey, err := base64.StdEncoding.DecodeString(publicKey)
	if err != nil {
		log.Println(err.Error())
		return "", fmt.Errorf("failed to decode key, error: %s", err.Error())
	}

	key, err := jwt.ParseRSAPublicKeyFromPEM(decodedKey)

	if err != nil {
		log.Println(err.Error())
		return "", fmt.Errorf("failed to parse decoded key, error: %s", err.Error())
	}

	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
			log.Println(err.Error())
			return nil, fmt.Errorf("unexpected method: %s", t.Header["alg"])
		}
		return key, nil
	})

	if err != nil {
		log.Println(err.Error())
		return "", fmt.Errorf("error occured: %s", err.Error())
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)

	if !ok || !parsedToken.Valid {
		log.Println(err.Error())
		return "", fmt.Errorf("error occured: %s", err.Error())
	}

	return claims["sub"], nil
}
