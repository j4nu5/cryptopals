package main

import (
	"fmt"
	"log"

	"github.com/j4nu5/cryptopals/solutions"
)

func main() {
	pt := "YELLOW SUBMARINE"
	padded, err := solutions.Pkcs7Pad([]byte(pt), 20)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(padded)
}
