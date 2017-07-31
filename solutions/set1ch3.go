package solutions

import (
	"encoding/hex"
	"math"
)

// CrackSingleByteXorCipher cracks a cipher where a single, unique byte
// is used to pad the entire byte array. Takes input as a hex encoded string.
func CrackSingleByteXorCipher(s string) (string, error) {
	bytes, err := hex.DecodeString(s)
	if err != nil {
		return "", err
	}

	mask := make([]byte, len(bytes))
	plainText := ""
	maxScore := 0
	for val := 0; val <= math.MaxUint8; val++ {
		for i := range mask {
			mask[i] = byte(val)
		}

		pt := string(Xor(bytes, mask))
		score := score(pt)
		if score > maxScore {
			maxScore = score
			plainText = pt
		}
	}

	return plainText, nil
}

func score(s string) int {
	score := 0
	for _, ch := range []byte{'e', 't', 'a', 'o', 'i', 'n', ' ', 's', 'h', 'r', 'd', 'l', 'u'} {
		score += frequency(s, ch)
	}

	return score
}

func frequency(s string, b byte) int {
	f := 0

	for _, ch := range []byte(s) {
		if toLower(ch) == b {
			f++
		}
	}

	return f
}

func toLower(b byte) byte {
	if b >= 'A' && b <= 'Z' {
		return b - 'A' + 'a'
	}

	return b
}
