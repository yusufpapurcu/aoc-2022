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

	res := 0
	for {
		line, _, err := r.ReadLine()
		if err != nil {
			log.Println(err)

			break
		}

		first := line[:len(line)/2]
		second := line[len(line)/2:]

		seen := map[byte]bool{}
		for _, a := range first {
			seen[a] = true
		}

		for _, a := range second {
			if _, ok := seen[a]; ok {
				seen[a] = false
			}
		}

		for item, isShared := range seen {
			if !isShared {
				res += getItemPoint(item)
			}
		}
	}

	return fmt.Sprint(res)
}

func solution2(f io.Reader) string {
	r := bufio.NewReader(f)

	res := 0
GLOB:
	for {
		var lines [3][]byte
		for i := 0; i < 3; i++ {
			line, _, err := r.ReadLine()
			if err != nil {
				log.Println(err)

				break GLOB
			}
			lines[i] = line
		}

		seenCount := map[byte]int{}
		for _, line := range lines {
			isSeenInLine := map[byte]bool{}
			for _, a := range line {
				if !isSeenInLine[a] {
					seenCount[a]++
					isSeenInLine[a] = true
				}
			}
		}

		for item, count := range seenCount {
			if count == 3 {
				res += getItemPoint(item)
			}
		}
	}

	return fmt.Sprint(res)
}

func getItemPoint(item byte) int {
	if int(item) > 64 && int(item) < 91 {
		return int(item) - 64 + 26
	}
	if int(item) > 96 && int(item) < 123 {
		return int(item) - 96
	}
	return 0
}
