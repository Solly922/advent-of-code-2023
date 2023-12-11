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

	m, dir := buildMap(scanner)
	fmt.Println(m, dir)
}

func buildMap(s *bufio.Scanner) (map[string][2]string, string) {
	m := make(map[string][2]string)
	var d string

	for s.Scan() {
		split := strings.Split(s.Text(), "=")
		if len(split) < 2 {
			if len(split) == 1 {
				d = split[0]
			}
			continue
		}

		node := split[0]
		options := strings.Split(split[1], ",")
		left := options[0][2:]
		right := options[1][1 : len(options[1])-1]

		m[node] = [2]string{left, right}
	}

	return m, d
}
