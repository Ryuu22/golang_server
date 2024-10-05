package utils

import (
	"strings"
)

func GetJWTFromAuthHeader(authHeader string) string {
	key, token := strings.Split(authHeader, " ")[0], strings.Split(authHeader, " ")[1]
	if key != "Bearer" {
		return ""
	}
	return token
}
