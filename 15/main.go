package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.ReadFile("test1.txt")
	if err != nil {
		panic(err)
	}

	codes := strings.Split(string(file), ",")

	sum := sumHashes(codes)
	fmt.Println(sum)

	placeInBoxes(codes)
}

func hashCode(code string) int {
	result := 0

	for _, c := range code {
		result += int(c)
		result *= 17
		result = result % 256
	}

	fmt.Println(result)
	return result
}

func sumHashes(codes []string) int {
	result := 0

	for _, code := range codes {
		result += hashCode(code)
	}

	return result
}

func placeInBoxes(codes []string) map[int][][2]string {
	result := make(map[int][][2]string)

	for _, code := range codes {
		split := strings.Split(code, "=")
		op := "="
		if len(split) == 1 {
			split = strings.Split(code, "-")
			op = "-"
		}

		boxNum := hashCode(split[0])
		fmt.Println(boxNum, op)
	}

	return result
}

func containsLabel(arr [][2]string, label string) (bool, int) {
	for i, a := range arr {
		if a[0] == label {
			return true, i
		}
	}

	return false, -1
}
