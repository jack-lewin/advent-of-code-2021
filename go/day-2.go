package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type movement struct {
	direction string
	amount int
}

func getMovement(s string) movement {
	line := strings.Split(s, " ")
	amount, err := strconv.Atoi(line[1])
	check(err)

	return movement {
		direction: line[0],
		amount: amount,
	}
}

func getInput() ([]movement, error) {
    file, err := os.Open("inputs/day-2.txt")
    check(err)
    defer file.Close()

    var lines []movement
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        lines = append(lines, getMovement(scanner.Text()))
    }
    return lines, scanner.Err()
}

func part1() int {
	input, err := getInput()
	check(err)

	dir := []int { 0, 0 }

	for i := 0; i < len(input); i++ {
		if input[i].direction == "forward" {
			dir[0] += input[i].amount
		} else if input[i].direction == "down" {
			dir[1] += input[i].amount
		} else if input[i].direction == "up" {
			dir[1] -= input[i].amount
		}
	}

	result := 1
	for _, v := range dir {
		result *= v
	}
	return result
}

func part2() int {
	input, err := getInput()
	check(err)

	dir := []int { 0, 0, 0 }

	for i := 0; i < len(input); i++ {
		if input[i].direction == "forward" {
			dir[0] += input[i].amount
			dir[1] += input[i].amount * dir[2]
		} else if input[i].direction == "down" {
			dir[2] += input[i].amount
		} else if input[i].direction == "up" {
			dir[2] -= input[i].amount
		}
	}

	return dir[0] * dir[1]
}

func main() {
	fmt.Println("part 1: ", part1())
	fmt.Println("part 2: ", part2())
}
