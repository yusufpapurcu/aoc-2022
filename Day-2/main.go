package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

var pointsStrategy1 = map[string]int{
	"A X": 1 + 3,
	"A Y": 2 + 6,
	"A Z": 3 + 0,

	"B X": 1 + 0,
	"B Y": 2 + 3,
	"B Z": 3 + 6,

	"C X": 1 + 6,
	"C Y": 2 + 0,
	"C Z": 3 + 3,
}

var pointsStrategy2 = map[string]int{
	"A X": 3 + 0,
	"A Y": 1 + 3,
	"A Z": 2 + 6,

	"B X": 1 + 0,
	"B Y": 2 + 3,
	"B Z": 3 + 6,

	"C X": 2 + 0,
	"C Y": 3 + 3,
	"C Z": 1 + 6,
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("File opening failed", err)
	}

	fmt.Println(solutionProblem1(f))

	f, err = os.Open("input.txt")
	if err != nil {
		log.Fatal("File opening failed", err)
	}

	fmt.Println(solutionProblem2(f))
}

func solutionProblem1(f io.Reader) string {
	r := bufio.NewReader(f)

	res := 0
	for {
		line, _, err := r.ReadLine()
		if err != nil {
			log.Println(err)

			break
		}

		point, ok := pointsStrategy1[string(line)]
		if !ok {
			fmt.Println(string(line))
		}
		res += point
	}

	return fmt.Sprint(res)
}

func solutionProblem2(f io.Reader) string {
	r := bufio.NewReader(f)

	res := 0
	for {
		line, _, err := r.ReadLine()
		if err != nil {
			log.Println(err)

			break
		}

		point, ok := pointsStrategy2[string(line)]
		if !ok {
			fmt.Println(string(line))
		}
		res += point
	}

	return fmt.Sprint(res)
}
