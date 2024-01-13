package hash

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
)

// HashPassword hashes a password using SHA-256 and returns the hashed value as a hexadecimal string.
func HashPassword(password string) string {
	hasher := sha256.New()

	_, err := io.WriteString(hasher, password)
	if err != nil {
		panic(err)
	}

	hashedBytes := hasher.Sum(nil)

	hashedPassword := hex.EncodeToString(hashedBytes)

	return hashedPassword
}

// ComparePasswords compares a hashed password with a plaintext password.
func ComparePasswords(hashedPassword, plaintextPassword string) bool {
	hashedAttempt := HashPassword(plaintextPassword)
	return hashedPassword == hashedAttempt
}
