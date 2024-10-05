package middleware

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"strings"

	"golang_server.dankbueno.com/internal/config"
	"golang_server.dankbueno.com/internal/models"
	"golang_server.dankbueno.com/internal/utils"
)

type JWTHeader struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}

type JWTPayload struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Iss      string `json:"iss"`
	Exp      string `json:"exp"`
}

func (payload JWTPayload) IsExpired() bool {
	return utils.IsExpired(payload.Exp)
}

type JWTSignature struct {
	Signature string `json:"signature"`
}

func GenerateJWTFromUser(user models.User) (string, error) {
	var header JWTHeader = JWTHeader{
		Alg: "HS256",
		Typ: "JWT",
	}

	var payload JWTPayload = JWTPayload{
		Id:       user.ID,
		Username: user.Username,
		Iss:      config.Issuer,
		Exp:      utils.GetExpirationTime(60),
	}

	headerJSON, err := json.Marshal(header)
	if err != nil {
		return "", err
	}

	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	unsignedToken := base64.URLEncoding.EncodeToString(headerJSON) + "." + base64.URLEncoding.EncodeToString(payloadJSON)

	h := hmac.New(sha256.New, []byte(config.JWTSecret))
	h.Write([]byte(unsignedToken))

	tokenString := unsignedToken + "." + base64.URLEncoding.EncodeToString(h.Sum(nil))

	return tokenString, nil
}

func VerifyJWT(tokenString string) (JWTPayload, error) {

	// Split token into parts
	tokenParts := strings.Split(tokenString, ".")
	if len(tokenParts) != 3 {
		return JWTPayload{}, errors.New("Invalid token")
	}

	// Verify signature
	unsignedToken := tokenParts[0] + "." + tokenParts[1]
	signature := tokenParts[2]

	h := hmac.New(sha256.New, []byte(config.JWTSecret))
	h.Write([]byte(unsignedToken))

	expectedSignature := base64.URLEncoding.EncodeToString(h.Sum(nil))

	if signature != expectedSignature {
		return JWTPayload{}, errors.New("Invalid token")
	}

	// Decode payload
	payloadJSON, err := base64.URLEncoding.DecodeString(tokenParts[1])
	if err != nil {
		return JWTPayload{}, err
	}

	var payload JWTPayload
	err = json.Unmarshal(payloadJSON, &payload)
	if err != nil {
		return JWTPayload{}, err
	}

	if payload.IsExpired() {
		return JWTPayload{}, errors.New("Token is expired")
	}

	if payload.Iss != config.Issuer {
		return JWTPayload{}, errors.New("Invalid token")
	}

	return payload, nil
}
