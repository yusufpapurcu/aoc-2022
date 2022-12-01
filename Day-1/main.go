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

	nextElf()
	for {
		line, _, err := r.ReadLine()
		if err != nil {
			log.Println(err)

			break
		}

		if isNewElf(string(line)) {
			nextElf()
			continue
		}

		if err := addToElf(string(line)); err != nil {
			log.Println(err)

			break
		}
	}

	res := findMostCalories()

	clearElfs()

	return fmt.Sprint(res)
}

func solutionProblem2(f io.Reader) string {
	r := bufio.NewReader(f)

	nextElf()
	for {
		line, _, err := r.ReadLine()
		if err != nil {
			log.Println(err)

			break
		}

		if isNewElf(string(line)) {
			nextElf()
			continue
		}

		if err := addToElf(string(line)); err != nil {
			log.Println(err)

			break
		}
	}

	res := findTop3Calories()

	clearElfs()

	return fmt.Sprint(res)
}
