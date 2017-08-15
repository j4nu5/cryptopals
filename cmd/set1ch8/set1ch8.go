package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

const filename = "/home/sinhak/go/src/github.com/j4nu5/cryptopals/cmd/set1ch8/8.txt"
const blockSize = 16 * 2

type pair struct {
	freq int
	ct   string
}
type pairs []pair

func (p pairs) Len() int {
	return len(p)
}

func (p pairs) Less(i, j int) bool {
	return p[i].freq < p[j].freq
}

func (p pairs) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func main() {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	s := bufio.NewScanner(file)
	uniqBlocks := make(map[string]int)
	for s.Scan() {
		ct := s.Text()
		uniqBlocks[ct] = calcUniqBlocks(ct)
	}

	printBlockInfo(uniqBlocks)
}

func calcUniqBlocks(s string) int {
	blockSet := make(map[string]bool)

	for len(s) > 0 {
		blockSet[s[:blockSize]] = true
		s = s[blockSize:]
	}

	return len(blockSet)
}

func printBlockInfo(blocks map[string]int) {
	var data pairs
	for k, v := range blocks {
		data = append(data, pair{v, k})
	}
	sort.Sort(data)

	for _, d := range data {
		fmt.Println(d.ct[:8], d.freq)
	}
}
