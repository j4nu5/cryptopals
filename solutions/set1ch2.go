package solutions

import (
	"encoding/hex"
)

// XorHexStrings returns the xor of two hex encoded strings.
func XorHexStrings(a, b string) (string, error) {
	aBytes, err := hex.DecodeString(a)
	if err != nil {
		return "", err
	}

	bBytes, err := hex.DecodeString(b)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(Xor(aBytes, bBytes)), nil
}

// Xor returns the xor of two byte arrays.
// These byte arrays can be of different lengths and will be aligned on the
// basis of their MSB.
func Xor(a, b []byte) []byte {
	c := make([]byte, max(len(a), len(b)))

	aIdx := len(a) - 1
	bIdx := len(b) - 1

	for cIdx := len(c) - 1; cIdx >= 0; cIdx-- {
		var aByte, bByte byte

		if aIdx >= 0 {
			aByte = a[aIdx]
		}
		if bIdx >= 0 {
			bByte = b[bIdx]
		}
		aIdx--
		bIdx--

		c[cIdx] = aByte ^ bByte
	}

	return c
}

func max(a, b int) int {
	if a >= b {
		return a
	}

	return b
}
