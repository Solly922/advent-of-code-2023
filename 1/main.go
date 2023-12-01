package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		digits := getDigits(scanner.Text())
		fmt.Println(digits)
	}
}

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
