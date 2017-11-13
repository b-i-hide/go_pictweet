package models

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

func HashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	return string(hash)
}
