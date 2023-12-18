package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	start := time.Now()
	codes := strings.Split(string(file), ",")

	// sum := sumHashes(codes)
	// fmt.Println(sum)

	boxes := placeInBoxes(codes)
	// fmt.Println(boxes)

	focusingPower := calcTotalFocusPower(boxes)
	end := time.Now()
	fmt.Println(focusingPower)

	fmt.Println("Total time: ", end.Sub(start))
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
		// fmt.Println(boxNum, op)

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
			} else {
				if op == "=" {
					v = append(v, [2]string{split[0], split[1]})
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

func calcTotalFocusPower(m map[int][][2]string) int {
	result := 0

	for k, v := range m {
		power := calcBoxFocusPower(k, v)
		result += power
	}

	return result
}

func calcBoxFocusPower(box int, arr [][2]string) int {
	result := 0
	if len(arr) < 1 {
		return result
	}

	for i, v := range arr {
		power := box + 1
		strength, err := strconv.Atoi(v[1])
		if err != nil {
			// fmt.Println(box, i, v)
			panic(err)
		}
		power *= ((i + 1) * strength)

		result += power
	}

	return result
}
