package main

import (
	"fmt"
	"strconv"

	advent "github.com/BenJuan26/advent2022"
)

func Part1() {
	lines, err := advent.ReadInput()
	if err != nil {
		panic(err)
	}

	currentElf := 0
	maxElf := 0
	for _, line := range lines {
		if line == "" {
			if currentElf > maxElf {
				maxElf = currentElf
			}
			currentElf = 0
			continue
		}

		num, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		currentElf += num
	}

	fmt.Println(maxElf)
}

func main() {
	if advent.ShouldRunPart1() {
		Part1()
	}

	if advent.ShouldRunPart2() {
		Part2()
	}
}
