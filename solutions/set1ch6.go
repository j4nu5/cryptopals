package solutions

import (
	"bytes"
	"encoding/hex"
	"log"
)

// HammingDistance returns the hamming distance between two byte slices.
func HammingDistance(a, b []byte) int {
	var minLen, dist int
	if len(a) < len(b) {
		minLen = len(a)
		dist = len(b) - len(a)
	} else {
		minLen = len(b)
		dist = len(a) - len(b)
	}

	for i := 0; i < minLen; i++ {
		dist += hammingDistance(a[i], b[i])
	}

	return dist
}

func hammingDistance(a, b byte) int {
	x := a ^ b

	popcnt := 0
	for x != 0 {
		popcnt += int(x & byte(1))
		x >>= 1
	}

	return popcnt
}

// CrackRepeatingXorCipher cracks repeating key xor cipher.
func CrackRepeatingXorCipher(ct []byte, keySize int) string {
	log.Printf("Dividing ct of len: %d into blocks of size: %d to get %d blocks\n", len(ct), keySize, len(ct)/keySize)
	cipherBlock := createCipherBlock(ct, keySize)
	log.Printf("Created cipher block of size %d x %d\n", len(cipherBlock), len(cipherBlock[0]))
	cipherBlockTranspose := transpose(cipherBlock)
	log.Printf("Transposed cipher block of size %d x %d\n", len(cipherBlockTranspose), len(cipherBlockTranspose[0]))

	var key []byte
	for _, cBlock := range cipherBlockTranspose {
		pt, k, _ := CrackRepeatingXorCipherBytes(cBlock)
		log.Println("Cracked:", pt)
		key = append(key, k)
	}

	cipher := NewRepeatingXorCipher(key, bytes.NewReader(ct))
	buf := new(bytes.Buffer)
	buf.ReadFrom(cipher)
	ptBytes, _ := hex.DecodeString(buf.String())
	return string(ptBytes)
}

func createCipherBlock(buf []byte, bs int) [][]byte {
	cipherBlock := make([][]byte, 0)

	p := 0
	for p+bs <= len(buf) {
		cipherBlock = append(cipherBlock, buf[p:p+bs])
		p += bs
	}

	return cipherBlock
}

func transpose(buf [][]byte) [][]byte {
	rows := len(buf)
	cols := len(buf[0])

	tr := make([][]byte, cols)
	for i := range tr {
		tr[i] = make([]byte, rows)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			tr[j][i] = buf[i][j]
		}
	}

	return tr
}
