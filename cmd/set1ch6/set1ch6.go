package main

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"fmt"
	"os"

	"github.com/j4nu5/cryptopals/solutions"
)

const filename = "6.txt"

func main() {
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fileScanner := bufio.NewScanner(file)
	ctBuf := new(bytes.Buffer)
	for fileScanner.Scan() {
		ctBuf.WriteString(fileScanner.Text())
	}
	ctBase64 := ctBuf.String()
	ct, err := base64.StdEncoding.DecodeString(ctBase64)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Choose keysize
	// keySize := 2
	// hd := float64(len(ct))
	// for ks := 1; ks <= 64 && (ks*2 <= len(ct)); ks++ {
	// 	d := float64(solutions.HammingDistance(ct[:ks], ct[ks:ks*2])) / float64(ks)
	// 	log.Println(ks, d)
	// 	if d < hd {
	// 		hd = d
	// 		keySize = ks
	// 	}
	// }
	// log.Println("Choosing key size:", keySize)

	for keySize := 1; keySize <= 40; keySize++ {
		pt := solutions.CrackRepeatingXorCipher(ct, keySize)
		if solutions.IsSane([]byte(pt)) {
			fmt.Println("======== DECRYPTED TEXT ========")
			fmt.Println(pt)
			fmt.Println("======== DECRYPTED TEXT ========")
		}
	}
}
