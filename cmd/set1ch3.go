package main

import (
	"fmt"

	"github.com/j4nu5/cryptopals/solutions"
)

func main() {
	pt, _ := solutions.CrackSingleByteXorCipher("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")
	fmt.Println(pt)
}
