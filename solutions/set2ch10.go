package solutions

import (
	"crypto/aes"
	"errors"
	"log"
)

const (
	BlockSize = 16
)

func CbcEncrypt(pt, key []byte) ([]byte, error) {
	paddedPt, err := Pkcs7Pad(pt, BlockSize)
	if err != nil {
		return nil, err
	}
	ct := make([]byte, len(paddedPt))

	numBlocks := len(paddedPt) / BlockSize
	xorBlock := make([]byte, BlockSize)

	log.Printf("%d blocks to encrypt\n", numBlocks)
	for i := 0; i < numBlocks; i++ {
		start := i * BlockSize
		end := start + BlockSize

		xorPt := Xor(xorBlock, paddedPt[start:end])
		enc, err := EcbEncryptBlock(xorPt, key)
		if err != nil {
			return nil, err
		}
		copy(ct[start:end], enc)
		copy(xorBlock, ct[start:end])
	}

	return ct, nil
}

func CbcDecrypt(ct, key []byte) ([]byte, error) {
	if len(ct)%BlockSize != 0 {
		return nil, errors.New("Cipher text must be in block size")
	}

	pt := make([]byte, len(ct))
	numBlocks := len(ct) / BlockSize
	xorBlock := make([]byte, BlockSize)

	log.Printf("%d blocks to decrypt\n", numBlocks)
	for i := 0; i < numBlocks; i++ {
		start := i * BlockSize
		end := start + BlockSize
		dec, err := EcbDecryptBlock(ct[start:end], key)
		if err != nil {
			return nil, err
		}

		copy(pt[start:end], Xor(xorBlock, dec))
		xorBlock = ct[start:end]
	}

	return pt, nil
}

func EcbEncryptBlock(pt, key []byte) ([]byte, error) {
	if len(pt) != BlockSize {
		return nil, errors.New("Length of plain text must be equal to block size")
	}
	if len(key) != BlockSize {
		return nil, errors.New("Length of key must be equal to block size")
	}

	cipher, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	ct := make([]byte, BlockSize)
	cipher.Encrypt(ct, pt)

	return ct, nil
}

func EcbDecryptBlock(ct, key []byte) ([]byte, error) {
	if len(ct) != BlockSize {
		return nil, errors.New("Length of cipher text must be equal to block size")
	}
	if len(key) != BlockSize {
		return nil, errors.New("Length of key must be equal to block size")
	}

	cipher, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	pt := make([]byte, BlockSize)
	cipher.Decrypt(pt, ct)

	return pt, nil
}
