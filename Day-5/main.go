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
	initialRaw := []string{}
	for {
		line, _, err := r.ReadLine()
		if err != nil {
			log.Println(err)

			break
		}
		if string(line) == "" {
			break
		}

		a := strings.ReplaceAll(string(line), "    ", " [-]")
		a = strings.ReplaceAll(a, "[", " ")
		a = strings.ReplaceAll(a, "]", " ")
		initialRaw = append(initialRaw, a)
	}

	fmt.Println(strings.Join(initialRaw, "\n"))
	stacks := map[int]*Stack{}
	stackId := strings.Fields(initialRaw[len(initialRaw)-1])
	for a := range stackId {
		stacks[a] = NewStack()
	}
	for i := len(initialRaw) - 2; i >= 0; i-- {
		crates := strings.Fields(initialRaw[i])
		for idx, a := range crates {
			if a != "-" {
				stack := stacks[idx]
				stack.Push(a)
			}
		}
	}

	for {
		line, _, err := r.ReadLine()
		if err != nil {
			log.Println(err)

			break
		}
		a := strings.ReplaceAll(string(line), "move", "")
		a = strings.ReplaceAll(a, "from", "")
		a = strings.ReplaceAll(a, "to", "")

		commandStr := strings.Fields(a)
		command := convertInt(commandStr)
		for i := 0; i < command[0]; i++ {
			from := stacks[command[1]-1]
			val := from.Pop()
			if val == "" {
				break
			}
			to := stacks[command[2]-1]
			to.Push(val)
		}
	}

	for i := 0; i < len(stacks); i++ {
		stack := stacks[i]
		fmt.Print(stack.Peek())
	}

	fmt.Println()

	return fmt.Sprint(stacks)
}

func solution2(f io.Reader) string {
	r := bufio.NewReader(f)
	initialRaw := []string{}
	for {
		line, _, err := r.ReadLine()
		if err != nil {
			log.Println(err)

			break
		}
		if string(line) == "" {
			break
		}

		a := strings.ReplaceAll(string(line), "    ", " [-]")
		a = strings.ReplaceAll(a, "[", " ")
		a = strings.ReplaceAll(a, "]", " ")
		initialRaw = append(initialRaw, a)
	}

	fmt.Println(strings.Join(initialRaw, "\n"))
	stacks := map[int]*Stack{}
	stackId := strings.Fields(initialRaw[len(initialRaw)-1])
	for a := range stackId {
		stacks[a] = NewStack()
	}
	for i := len(initialRaw) - 2; i >= 0; i-- {
		crates := strings.Fields(initialRaw[i])
		for idx, a := range crates {
			if a != "-" {
				stack := stacks[idx]
				stack.Push(a)
			}
		}
	}

	for {
		line, _, err := r.ReadLine()
		if err != nil {
			log.Println(err)

			break
		}
		a := strings.ReplaceAll(string(line), "move", "")
		a = strings.ReplaceAll(a, "from", "")
		a = strings.ReplaceAll(a, "to", "")

		commandStr := strings.Fields(a)
		command := convertInt(commandStr)
		toPush := []string{}
		for i := 0; i < command[0]; i++ {
			from := stacks[command[1]-1]
			val := from.Pop()
			if val == "" {
				break
			}
			toPush = append(toPush, val)
		}
		to := stacks[command[2]-1]
		for i := len(toPush) - 1; i >= 0; i-- {
			to.Push(toPush[i])
		}
	}

	for i := 0; i < len(stacks); i++ {
		stack := stacks[i]
		fmt.Print(stack.Peek())
	}

	fmt.Println()

	return fmt.Sprint(stacks)
}

func convertInt(a []string) [3]int {
	var nums [3]int
	nums[0], _ = strconv.Atoi(a[0])
	nums[1], _ = strconv.Atoi(a[1])
	nums[2], _ = strconv.Atoi(a[2])

	return nums
}
