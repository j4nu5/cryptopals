package solutions

import (
	"encoding/hex"
	"io"
)

// RepeatingXorCipher implements a repeating xor cipher.
type RepeatingXorCipher struct {
	key      []byte
	keyIndex int
	pt       io.Reader
}

func (c *RepeatingXorCipher) Read(p []byte) (int, error) {
	buf := make([]byte, len(p)/2)
	n, err := c.pt.Read(buf)

	for i := 0; i < n; i++ {
		buf[i] ^= c.key[c.keyIndex]
		c.keyIndex = (c.keyIndex + 1) % len(c.key)
	}

	hex.Encode(p, buf)
	return n * 2, err
}

// NewRepeatingXorCipher returns a new RepeatingXorCipher.
func NewRepeatingXorCipher(key []byte, pt io.Reader) *RepeatingXorCipher {
	return &RepeatingXorCipher{
		key:      key,
		keyIndex: 0,
		pt:       pt,
	}
}
