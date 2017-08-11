package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/j4nu5/cryptopals/solutions"
)

func main() {
	cts, err := getCiphers()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, ct := range cts {
		pt, _ := solutions.CrackSingleByteXorCipher(ct)
		fmt.Println("\nText: ", pt)
	}
}

func getCiphers() ([]string, error) {
	file, err := os.Open("4.txt")
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, nil
}
