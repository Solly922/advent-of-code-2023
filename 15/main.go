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

	boxes := placeInBoxes(codes)
	fmt.Println(boxes)
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

		v, ok := result[boxNum]
		if !ok {
			v = [][2]string{{split[0], split[1]}}
		} else {
			contains, i := containsLabel(v, split[0])
			if contains {
				if op == "=" {
					v = replaceLabel(v, i, split[0], split[1])
				} else {
					v = removeLabel(v, i, split[0])
				}
			}
		}
		result[boxNum] = v
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

func replaceLabel(arr [][2]string, idx int, label string, value string) [][2]string {
	arr[idx][1] = value
	return arr
}

func removeLabel(arr [][2]string, idx int, label string) [][2]string {
	arr = append(arr[:idx], arr[idx+1:]...)
	return arr
}
