package solutions

import (
	"encoding/hex"
	"math"
	"unicode"
)

// CrackSingleByteXorCipher cracks a cipher where a single, unique byte
// is used to pad the entire byte array. Takes input as a hex encoded string.
func CrackSingleByteXorCipher(s string) (string, error) {
	bytes, err := hex.DecodeString(s)
	if err != nil {
		return "", err
	}

	pt, _, err := CrackRepeatingXorCipherBytes(bytes)
	return pt, err
}

// CrackRepeatingXorCipherBytes cracks cipher text bytes.
func CrackRepeatingXorCipherBytes(ct []byte) (string, byte, error) {
	mask := make([]byte, len(ct))
	plainText := ""
	maxScore := 0
	var key byte
	for val := 0; val <= math.MaxUint8; val++ {
		for i := range mask {
			mask[i] = byte(val)
		}

		ptBytes := Xor(ct, mask)
		if !IsSane(ptBytes) {
			continue
		}

		score := score(ptBytes)
		if score > maxScore {
			maxScore = score
			plainText = string(ptBytes)
			key = byte(val)
		}
	}

	return plainText, key, nil
}

// IsSane returns whether the given byte array is a printable unicode string.
func IsSane(buf []byte) bool {
	for _, b := range buf {
		if !unicode.IsGraphic(rune(b)) && !unicode.IsSpace(rune(b)) {
			return false
		}
	}

	return true
}

func score(buf []byte) int {
	score := 0

	for _, ch := range []byte{'e', 't', 'a', 'o', 'i', 'n', ' ', 's', 'h', 'r', 'd', 'l', 'u'} {
		score += frequency(buf, ch)
	}

	return score
}

func frequency(buf []byte, b byte) int {
	f := 0

	for _, ch := range buf {
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
