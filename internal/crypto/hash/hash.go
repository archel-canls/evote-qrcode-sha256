package hash

import (
	"crypto/sha256"
	"encoding/hex"
)

// Generate menghasilkan hash SHA256 dari input string
func Generate(input string) string {
	sum := sha256.Sum256([]byte(input))
	return hex.EncodeToString(sum[:])
}

// GenerateWithSalt menghasilkan SHA256 dengan tambahan salt (lebih aman)
func GenerateWithSalt(input, salt string) string {
	sum := sha256.Sum256([]byte(input + salt))
	return hex.EncodeToString(sum[:])
}
