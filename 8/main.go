package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("test1.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	m, ins, dest := buildMap(scanner)
	fmt.Println(m, ins)
	steps := countSteps(m, ins, dest)
	fmt.Println(steps)
}

func buildMap(s *bufio.Scanner) (quest map[string][2]string, instructions string, destination string) {
	m := make(map[string][2]string)
	var d string
	var dest string

	for s.Scan() {
		split := strings.Split(s.Text(), "=")
		if len(split) < 2 {
			fmt.Println("Split?", split)
			if len(split) == 1 && split[0] != "" {
				d = split[0]
			}
			continue
		}

		node := strings.Trim(split[0], " ")
		options := strings.Split(split[1], ",")
		left := options[0][2:]
		right := options[1][1 : len(options[1])-1]

		m[node] = [2]string{left, right}
		dest = node
	}

	return m, d, dest
}

func countSteps(m map[string][2]string, ins string, dest string) int {
	end := false
	steps := 0
	node := "AAA"
	if ins == "" {
		fmt.Println("No instructions")
		return 0
	}
	fmt.Println(m)

	for !end {
		if steps > 90 {
			fmt.Println("Too many steps")
			end = true
		}
		for _, v := range ins {
			// fmt.Println(node, m[node])
			if node == dest {
				return steps
			}

			if string(v) == "L" {
				node = m[node][0]
			}
			if string(v) == "R" {
				node = m[node][1]
			}
			steps++
		}
	}

	return steps
}
