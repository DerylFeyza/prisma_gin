package bcrypt

import (
	"os"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword hashes the given password with the specified cost (number of rounds).
func HashPassword(password string) (string, error) {
	rounds, err := strconv.Atoi(os.Getenv("BCRYPT_ROUNDS"))
	if err != nil {
		return "", err
	}
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), rounds)
	return string(bytes), err
}
