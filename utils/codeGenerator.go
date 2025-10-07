package utils

import (
	"crypto/rand"
	"math/big"
)

// GenerateShortCode creates a random alphanumeric short code of given length.
func GenerateShortCode(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	code := make([]byte, length)

	for i := range code {
		// crypto/rand ensures secure randomness
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			// fallback if something goes wrong
			code[i] = charset[0]
			continue
		}
		code[i] = charset[num.Int64()]
	}

	return string(code)
}
