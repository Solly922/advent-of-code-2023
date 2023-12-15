package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	codes := strings.Split(string(file), ",")

	sum := sumHashes(codes)
	fmt.Println(sum)
}

func hashCode(code string) int {
	result := 0

	for _, c := range code {
		result += int(c)
		result *= 17
		result = result % 256
	}

	return result
}

func sumHashes(codes []string) int {
	result := 0

	for _, code := range codes {
		result += hashCode(code)
	}

	return result
}
