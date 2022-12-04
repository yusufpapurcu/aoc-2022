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
		lineRaw, _, err := r.ReadLine()
		if err != nil {
			log.Println(err)

			break
		}
		fields := strings.Fields(strings.Replace(strings.Replace(string(lineRaw), "-", " ", -1), ",", " ", -1))
		nums := convertInt(fields)

		front := nums[0] - nums[2]
		back := nums[1] - nums[3]

		if (front >= 0 && back <= 0) || (front <= 0 && back >= 0) {
			res++
		}
	}
	return fmt.Sprint(res)
}

func solution2(f io.Reader) string {
	r := bufio.NewReader(f)
	res := 0
	for {
		lineRaw, _, err := r.ReadLine()
		if err != nil {
			log.Println(err)

			break
		}
		fields := strings.Fields(strings.Replace(strings.Replace(string(lineRaw), "-", " ", -1), ",", " ", -1))
		nums := convertInt(fields)

		if nums[1]-nums[2] >= 0 && nums[0]-nums[3] <= 0 {
			res++
		}
	}
	return fmt.Sprint(res)
}

func convertInt(a []string) [4]int {
	var nums [4]int
	nums[0], _ = strconv.Atoi(a[0])
	nums[1], _ = strconv.Atoi(a[1])
	nums[2], _ = strconv.Atoi(a[2])
	nums[3], _ = strconv.Atoi(a[3])

	return nums
}
