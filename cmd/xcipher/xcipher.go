package main

import (
	"fmt"
	"io"
	"os"

	"github.com/j4nu5/cryptopals/solutions"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Enter the encryption key in ASCII as the only argument")
		os.Exit(1)
	}

	key := os.Args[1]

	cipher := solutions.NewRepeatingXorCipher([]byte(key), os.Stdin)
	io.Copy(os.Stdout, cipher)
}
