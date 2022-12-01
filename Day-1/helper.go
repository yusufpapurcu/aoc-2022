package main

import (
	"fmt"
	"strconv"
)

var elfs = []int{}

func isNewElf(line string) bool {
	return line == ""
}

func addToElf(line string) error {
	a, err := strconv.Atoi(line)
	if err != nil {
		return err
	}
	elfs[len(elfs)-1] += a
	return nil
}

func nextElf() {
	elfs = append(elfs, 0)
}

func clearElfs() {
	elfs = []int{}
}

func findMostCalories() int {
	res := 0
	for _, a := range elfs {
		if a > res {
			res = a
		}
	}
	return res
}

func findTop3Calories() int {
	first := 0
	second := 0
	third := 0

	for _, a := range elfs {
		if a >= first {
			third = second
			second = first
			first = a
			continue
		}
		if a >= second {
			third = second
			second = a
			continue
		}
		if a >= third {
			third = a
			continue
		}
	}

	fmt.Println(first, second, third)
	return first + second + third
}
