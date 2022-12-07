package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(solution1(f))

	f, err = os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(solution2(f))
}

func solution1(f io.Reader) string {
	r := bufio.NewReader(f)
GOBO:
	for {
		line, _, err := r.ReadLine()
		if err != nil {
			log.Println(err)

			break
		}
		for a := range line {
			if a+4 > len(line) {
				break
			}

			if unique4byte(line[a : a+4]) {
				fmt.Println(a + 4)
				break GOBO
			}
		}
	}
	return ""
}

func solution2(f io.Reader) string {
	r := bufio.NewReader(f)
	start := 0

	for {
		line, _, err := r.ReadLine()
		if err != nil {
			log.Println(err)

			break
		}
		for a := range line {
			if a+4 > len(line) {
				break
			}

			if unique4byte(line[a : a+4]) {
				start = a + 4
				break
			}
		}
		msgLine := line[start:]
		for a := range msgLine {
			if a+14 > len(msgLine) {
				break
			}

			if unique4byte(msgLine[a : a+14]) {
				fmt.Println(a + 14 + start)
				break
			}
		}
	}
	return ""
}

func unique4byte(bytes []byte) bool {
	for a := range bytes {
		for b := range bytes {
			if a == b {
				continue
			}
			if bytes[a] == bytes[b] {
				return false
			}
		}
	}
	return true
}
