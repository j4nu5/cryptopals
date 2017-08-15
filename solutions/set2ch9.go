package solutions

import (
	"errors"
	"math"
	"strconv"
)

// Pkcs7Pad pads the `plain text' b to a multiple of block size bs.
func Pkcs7Pad(b []byte, bs int) ([]byte, error) {
	if bs <= 0 {
		return nil, errors.New("Block size must be greater than 0")
	}
	if bs > math.MaxUint8 {
		return nil, errors.New("Block size must be less than" + strconv.Itoa(math.MaxUint8))
	}

	p := make([]byte, len(b))
	copy(p, b)

	padLen := bs - len(p)%bs
	for i := 0; i < padLen; i++ {
		p = append(p, byte(padLen))
	}

	return p, nil
}
