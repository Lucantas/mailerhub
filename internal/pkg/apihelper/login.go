package apihelper

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func SecurePassword(pwd string) string {
	hash, err := bcrypt.GenerateFromPassword(
		[]byte(pwd),
		bcrypt.DefaultCost,
	)

	if err != nil {
		// TODO: handle this error in a propely way
		log.Println("Error generating hash: ", err)
	}

	return string(hash)
}

func PasswordMatch(hashedPwd string, plainPwd string) bool {

	err := bcrypt.CompareHashAndPassword(
		[]byte(hashedPwd),
		[]byte(plainPwd),
	)

	if err != nil {
		log.Println("Error while checking if password match: ", err)
		return false
	}

	return true
}
