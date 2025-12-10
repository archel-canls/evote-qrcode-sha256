package crypto

import (
	"crypto/sha256"
	"encoding/hex"
)

// GenerateHash returns SHA-256 hash of input string
func GenerateHash(input string) string {
	hash := sha256.Sum256([]byte(input))
	return hex.EncodeToString(hash[:])
}
