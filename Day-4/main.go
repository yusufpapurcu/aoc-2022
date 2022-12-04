package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
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
		fields := strings.Split(string(line), ",")
		p1 := strings.Split(fields[0], "-")
		p2 := strings.Split(fields[1], "-")
		p1a1 := convertInt(p1[0])
		p1a2 := convertInt(p1[1])
		p2a1 := convertInt(p2[0])
		p2a2 := convertInt(p2[1])

		front := p1a1 - p2a1
		back := p1a2 - p2a2

		if front >= 0 && back <= 0 {
			res++
		} else if front <= 0 && back >= 0 {
			res++
		}
	}
	return fmt.Sprint(res)
}

func solution2(f io.Reader) string {
	r := bufio.NewReader(f)
	res := 0
	for {
		line, _, err := r.ReadLine()
		if err != nil {
			log.Println(err)

			break
		}
		fields := strings.Split(string(line), ",")
		p1 := strings.Split(fields[0], "-")
		p2 := strings.Split(fields[1], "-")
		p1a1 := convertInt(p1[0])
		p1a2 := convertInt(p1[1])
		p2a1 := convertInt(p2[0])
		p2a2 := convertInt(p2[1])

		if p1a2-p2a1 >= 0 {
			if p1a1-p2a2 <= 0 {
				res++
			}
		}

		// .....67..  6-7
		// ...45....  4-5

		// front := p1a1 - p2a1
		// back := p1a2 - p2a2

		// if front >= 0 && back <= 0 {
		// 	res++
		// } else if front <= 0 && back >= 0 {
		// 	res++
		// }
	}
	return fmt.Sprint(res)
}

func convertInt(a string) int {
	res, _ := strconv.Atoi(a)
	return res
}
