package main

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"fmt"
	"log"
	"os"

	"github.com/j4nu5/cryptopals/solutions"
)

const (
	key      = "YELLOW SUBMARINE"
	filename = "/home/sinhak/go/src/github.com/j4nu5/cryptopals/cmd/set2ch10/10.txt"
)

func main() {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}

	s := bufio.NewScanner(file)
	var ctBuffer bytes.Buffer
	for s.Scan() {
		ctBuffer.WriteString(s.Text())
	}

	ct := ctBuffer.String()
	ctBytes, err := base64.StdEncoding.DecodeString(ct)
	if err != nil {
		log.Fatal(err)
	}

	ptBytes, err := solutions.CbcDecrypt(ctBytes, []byte(key))
	if err != nil {
		log.Fatal(err)
	}

	paddingBytes := int(ptBytes[len(ptBytes)-1])
	fmt.Println(string(ptBytes[:len(ptBytes)-paddingBytes]))
}
