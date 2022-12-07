package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(solution1(f))
}

func solution1(f io.Reader) string {
	r := bufio.NewReader(f)
	dirSize := map[string]int{}
	dirCalcAfter := map[string][]string{}
	dirCalcAfterJustNames := []string{}

	currentDir := ""
	for {
		l, _, err := r.ReadLine()
		if err != nil {
			log.Println(err)

			break
		}
		line := string(l)

		if strings.HasPrefix(line, "$") {
			if line == "$ ls" {
				for {
					l, _, err := r.ReadLine()
					if err != nil {
						log.Println(err)

						break
					}
					line = string(l)
					if strings.HasPrefix(line, "$") {
						currentDir = dirSelection(currentDir, line)
						break
					}

					details := strings.Fields(line)
					if details[0] == "dir" {
						dirCalcAfter[currentDir] = append(dirCalcAfter[currentDir], currentDir+details[1]+"/")
						dirCalcAfterJustNames = append(dirCalcAfterJustNames, currentDir)
					} else {
						a, _ := strconv.Atoi(details[0])
						dirSize[currentDir] += a
					}
				}
			} else {
				currentDir = dirSelection(currentDir, line)
			}
		}
	}

	sort.Slice(dirCalcAfterJustNames, func(i, j int) bool {
		return dirCalcAfterJustNames[i] > dirCalcAfterJustNames[j]
	})

	for _, name := range dirCalcAfterJustNames {
		dirs, ok := dirCalcAfter[name]
		if !ok {
			continue
		}
		for _, subdirs := range dirs {
			dirSize[name] += dirSize[subdirs]
		}
		delete(dirCalcAfter, name)
	}
	res1 := 0
	neededSpace := 30000000 - (70000000 - dirSize["/"])
	res2 := 30000000

	for key, val := range dirSize {
		if val < 100000 {
			fmt.Println(key)
			res1 += val
		}
		if val >= neededSpace {
			if val < res2 {
				res2 = val
			}
		}
	}
	return fmt.Sprint(res1, res2)
}

func dirSelection(currentDir, line string) string {
	params := strings.Fields(line)
	if params[2] == ".." {
		path := strings.Split(currentDir, "/")
		currentDir = strings.Join(path[:len(path)-2], "/") + "/"
	} else if params[2] == "/" {
		currentDir = "/"
	} else {
		currentDir += params[2] + "/"
	}

	return currentDir
}
