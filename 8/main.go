package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	m, ins, start, dest := buildMap(scanner)
	startTime := time.Now()
	// fmt.Println(m, ins)
	steps := countSteps(m, ins, start, dest)
	endTime := time.Now()
	fmt.Println("Time taken:", endTime.Sub(startTime))
	fmt.Println(steps)

	startTime = time.Now()
	aNodes := getANodes(&m)
	allSteps := countAllSteps(&m, ins, aNodes)
	lcm := lowestCommonMultiple(allSteps)
	endTime = time.Now()
	fmt.Println(lcm)
	fmt.Println("Part 2 Time taken:", endTime.Sub(startTime))
}

func buildMap(s *bufio.Scanner) (quest map[string][2]string, instructions string, start string, destination string) {
	m := make(map[string][2]string)
	var d string
	var dest string

	for s.Scan() {
		split := strings.Split(s.Text(), "=")
		if len(split) < 2 {
			if len(split) == 1 && split[0] != "" {
				d = split[0]
			}
			continue
		}

		node := strings.Trim(split[0], " ")
		if start == "" {
			start = node
		}
		options := strings.Split(split[1], ",")
		left := options[0][2:]
		right := options[1][1 : len(options[1])-1]

		m[node] = [2]string{left, right}
		dest = node
	}

	return m, d, start, dest
}

func countSteps(m map[string][2]string, ins string, start string, dest string) int {
	end := false
	steps := 0
	node := "AAA"
	if ins == "" {
		fmt.Println("No instructions")
		return 0
	}
	fmt.Println(node, dest)

	for !end {
		if steps > 999999 {
			fmt.Println("Too many steps in countSteps")
			end = true
		}
		for _, v := range ins {
			// fmt.Println(node, string(v))
			if node == "ZZZ" {
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

func getANodes(m *map[string][2]string) []string {
	var result []string

	for k := range *m {
		if k[2] == 'A' {
			result = append(result, k)
		}
	}

	return result
}

func countStepsModif(m *map[string][2]string, ins string, start string) int {
	end := false
	steps := 0
	node := start
	if ins == "" {
		fmt.Println("No instructions")
		return 0
	}

	for !end {
		if steps > 999999 {
			fmt.Println("Too many steps in countStepsModif")
			end = true
		}
		for _, v := range ins {
			// fmt.Println(node, string(v))
			if node[2] == 'Z' {
				return steps
			}

			if string(v) == "L" {
				node = (*m)[node][0]
			}
			if string(v) == "R" {
				node = (*m)[node][1]
			}
			steps++
		}
	}

	return steps
}

func countAllSteps(m *map[string][2]string, ins string, aNodes []string) []int {
	var result []int

	for _, v := range aNodes {
		result = append(result, countStepsModif(m, ins, v))
	}

	return result
}

func gcd(a, b int) int {
	if a == 0 {
		return b
	}

	return gcd(b%a, a)
}

func lowestCommonMultiple(nums []int) int {
	result := nums[0]

	for i := 1; i < len(nums); i++ {
		result = (result * nums[i]) / gcd(result, nums[i])
	}

	return result
}
