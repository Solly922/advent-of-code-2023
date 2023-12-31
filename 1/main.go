package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type void struct{}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	sum := 0
	timeStart := time.Now()
	for scanner.Scan() {
		// digits := getDigits(scanner.Text())
		// fmt.Println("-----")
		digits := getValidDigits(scanner.Text())
		// fmt.Println(digits)
		sum += digits
	}

	fmt.Println(sum)
	timeEnd := time.Now()
	fmt.Println("MINE:", timeEnd.Sub(timeStart))

	timeStart = time.Now()
	otherSolution()
	timeEnd = time.Now()
	fmt.Println("OTHER:", timeEnd.Sub(timeStart))
}

func otherSolution() {
	input, _ := os.ReadFile("input.txt")

	calc := func(r *strings.Replacer) (result int) {
		for _, s := range strings.Fields(string(input)) {
			s = r.Replace(r.Replace(s))
			result += 10 * int(s[strings.IndexAny(s, "123456789")]-'0')
			result += int(s[strings.LastIndexAny(s, "123456789")] - '0')
		}
		return
	}

	fmt.Println(calc(strings.NewReplacer()))
	fmt.Println(calc(strings.NewReplacer("one", "o1e", "two", "t2o", "three", "t3e", "four",
		"f4r", "five", "f5e", "six", "s6x", "seven", "s7n", "eight", "e8t", "nine", "n9e")))
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
	var resultString string
	arr := getDigitsArray(s)
	resultString += strconv.Itoa(arr[0])
	resultString += strconv.Itoa(arr[len(arr)-1])

	result, err := strconv.Atoi(resultString)
	if err != nil {
		panic(err)
	}
	return result
}

func getDigitsArray(s string) []int {
	var result []int
	var v void
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
	left, right := 0, 0
	// fmt.Println(s)

	for left < len(s) {
		letter := s[left]
		var currString string
		// fmt.Println("result", result)
		// fmt.Println("left", string(letter))
		isDigit := letter >= '0' && letter <= '9'
		if isDigit {
			digit, err := strconv.Atoi(string(letter))
			if err != nil {
				panic(err)
			}
			result = append(result, digit)
			left++
			right = left
			continue
		}

		set := make(map[string]void) // set of possible numbers
		// fill set
		for key := range spelledNumbers {
			set[key] = v
		}

		for right < len(s) {
			letter := s[right]
			position := right - left
			currString += string(letter)

			if _, ok := set[currString]; ok {
				result = append(result, spelledNumbers[currString])
				break
			}

			for key := range set {
				if position > len(key)-1 {
					delete(set, key)
					continue
				}

				if key[position] != letter {
					delete(set, key)
				}
			}

			if len(set) == 0 {
				break
			}

			right++
		}

		left++
		right = left
	}

	// fmt.Println(result)
	return result
}
