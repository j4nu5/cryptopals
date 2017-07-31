package solutions

import (
	"encoding/base64"
	"encoding/hex"
)

// HexToBase64 converts a base 16 encoded string to a
// base 64(standard alphabet) string. Optionally it returns an error.
func HexToBase64(s string) (string, error) {
	bytes, err := hex.DecodeString(s)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(bytes), nil
}
