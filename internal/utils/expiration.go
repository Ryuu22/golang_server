package utils

import (
	"log"
	"time"
)

func GetExpirationTime(seconds int) string {
	return time.Now().Add(time.Duration(seconds) * time.Second).Format(time.RFC3339)
}

func IsExpired(expiration string) bool {
	expirationTime, err := time.Parse(time.RFC3339, expiration)
	log.Println("Expiration: ", expiration)
	log.Println("Expiration time: ", expirationTime)
	if err != nil {
		return true
	}
	return time.Now().After(expirationTime)
}
