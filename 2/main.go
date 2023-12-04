package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	sum := 0

	for scanner.Scan() {
		ok, id := isPossibleGame(scanner.Text())
		if ok {
			sum += id
		}
	}

	fmt.Println(sum)
}

func isPossibleGame(s string) (bool, int) {
	gameId := strings.Split(strings.Split(s, ":")[0], " ")[1]
	fmt.Println(gameId)

	maxes := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	sets := strings.Split(strings.Split(s, ":")[1], ";")

	for _, set := range sets {
		cubes := strings.Split(set, ",")

		for _, cube := range cubes {
			cube = strings.Trim(cube, " ")
			num, err := strconv.Atoi(strings.Split(cube, " ")[0])
			if err != nil {
				panic(err)
			}
			color := strings.Split(cube, " ")[1]

			if num > maxes[color] {
				return false, 0
			}
		}
	}

	id, err := strconv.Atoi(gameId)
	if err != nil {
		panic(err)
	}

	return true, id
}
