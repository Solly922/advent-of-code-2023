package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type void struct{}

func main() {
	file, err := os.Open("practice.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	sum := 0

	for scanner.Scan() {
		// digits := getDigits(scanner.Text())
		digits := getValidDigits(scanner.Text())
		fmt.Println(digits)
		sum += digits
	}

	fmt.Println(sum)
}

// first solution
// finds the first and last digit of a string
func getDigits(s string) int {
	var result string
	stoppedAt := 0
	for i := 0; i < len(s); i++ {
		letter := s[i]
		isDigit := letter >= '0' && letter <= '9'

		if isDigit {
			result += string(letter)
			stoppedAt = i
			break
		}
	}
	for i := len(s) - 1; i >= stoppedAt; i-- {
		letter := s[i]
		isDigit := letter >= '0' && letter <= '9'

		if isDigit {
			result += string(letter)
			break
		}
	}

	digits, err := strconv.Atoi(result)
	if err != nil {
		panic(err)
	}

	return digits
}

// second solution
// finds the first and last digit of a string, but spelled out numbers count
func getValidDigits(s string) int {
	var v void
	var resultString string
	var spelledNumbers = map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}
	keys := make([]string, len(spelledNumbers))
	for k, val := range spelledNumbers {
		keys[val-1] = k
	}
	left, right := 0, 0

	for left < len(s) {
		letter := s[left]
		isDigit := letter >= '0' && letter <= '9'
		if isDigit {
			resultString += string(letter)
			break
		}

		set := make(map[string]void)
		for right < len(s) {
			letter := s[right]
			position := right - left

			if len(set) == 1 {
				ok := false
				var val int
				for key := range set {
					val, ok = spelledNumbers[key]
					break
				}

				if ok {
					resultString += strconv.Itoa(val)
					break
				}
			}

			for _, key := range keys {
				if position > len(key)-1 {
					delete(set, key)
					continue
				}

				if key[position] == letter {
					set[key] = v
				} else {
					delete(set, key)
				}
			}

			if len(set) == 0 {
				break
			}

			right++
		}
		if len(resultString) > 0 {
			break
		}

		left++
		right = left
	}

	left, right = len(s)-1, len(s)-1
	for left >= 0 {
		letter := s[left]
		isDigit := letter >= '0' && letter <= '9'
		if isDigit {
			resultString += string(letter)
			break
		}

		set := make(map[string]void)
		for right >= 0 {
			letter := s[right]
			position := right - left

			if len(set) == 1 {
				ok := false
				var val int
				for key := range set {
					val, ok = spelledNumbers[key]
					break
				}

				if ok {
					resultString += strconv.Itoa(val)
					break
				}
			}

			for _, key := range keys {
				if position > len(key)-1 {
					delete(set, key)
					continue
				}

				if key[position] == letter {
					set[key] = v
				} else {
					delete(set, key)
				}
			}

			if len(set) == 0 {
				break
			}

			right--
		}
		if len(resultString) > 1 {
			break
		}

		left--
		right = left
	}

	digits, err := strconv.Atoi(resultString)
	if err != nil {
		panic(err)
	}

	return digits
}
