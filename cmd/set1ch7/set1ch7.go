package main

import (
	"bufio"
	"bytes"
	"crypto/aes"
	"encoding/base64"
	"fmt"
	"log"
	"os"
)

const filename = "/home/sinhak/go/src/github.com/j4nu5/cryptopals/cmd/set1ch7/7.txt"
const key = "YELLOW SUBMARINE"

func main() {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	s := bufio.NewScanner(file)
	var buf bytes.Buffer
	for s.Scan() {
		buf.WriteString(s.Text())
	}

	ct, err := base64.StdEncoding.DecodeString(buf.String())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(decrypt(ct, []byte(key)))
}

func decrypt(ct []byte, key []byte) string {
	cipher, err := aes.NewCipher(key)
	if err != nil {
		log.Fatal(err)
	}

	pt := make([]byte, len(ct))
	bs := aes.BlockSize

	p := pt
	for len(ct) > 0 {
		cipher.Decrypt(p, ct)
		p = p[bs:]
		ct = ct[bs:]
	}

	return string(pt)
}
