package account

import (
	"io"
	"log"

	"crypto/rand"

	"golang.org/x/crypto/bcrypt"
)

var INTEGER_TABLE = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}

//EncodeToString is used to create a 6 digit verification code
func EncodeToString(max int) string {
	b := make([]byte, max)
	n, err := io.ReadAtLeast(rand.Reader, b, max)
	if n != max {
		panic(err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = INTEGER_TABLE[int(b[i])%len(INTEGER_TABLE)]
	}
	return string(b)
}

//HashPassword is used to encrypt the password before storage
func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	if err != nil {
		log.Panic(err)
	}

	return string(bytes)
}

//VerifyPassword checks the hashes of 2 passwords are equal
func VerifyPassword(userPassword string, providedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(providedPassword), []byte(userPassword))

	if err != nil {
		return false
	}

	return true
}
