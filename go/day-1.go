package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func strToIntSlice(s string) ([]int, error) {
	var result []int
	rawJson := "[" + strings.ReplaceAll(s, "\n", ",") + "]"
	err := json.Unmarshal([]byte(rawJson), &result)

	return result, err
}

func getInput() ([]int, error) {
	rawInput, err := os.ReadFile("inputs/day-1.txt")
	check(err)

	return strToIntSlice(string(rawInput))
}

func getRollingAverage(input []int) ([]int) {
	var result []int

	for i := 0; i < len(input) - 2; i++ {
		result = append(result, input[i] + input[i + 1] + input[i + 2])
	}

	return result
}

func part1() int {
	input, err := getInput()
	check(err)

	count := 0

	for i := 1; i < len(input); i++ {
		if input[i] > input[i - 1] {
			count++
		}
	}

	return count
}

func part2() int {
	input, err := getInput()
	check(err)
	rollingAverage := getRollingAverage(input)

	count := 0

	for i := 1; i < len(rollingAverage); i++ {
		if rollingAverage[i] > rollingAverage[i - 1] {
			count++
		}
	}

	return count
}

func main() {
	fmt.Println("part 1: ", part1())
	fmt.Println("part 2: ", part2())
}
