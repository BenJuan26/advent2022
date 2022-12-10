package main

import (
	"fmt"
	"strconv"
	"strings"

	advent "github.com/BenJuan26/advent2022"
)

func updateStrength(x, currentCycle int) int {
	if currentCycle == 20 ||
		currentCycle == 60 ||
		currentCycle == 100 ||
		currentCycle == 140 ||
		currentCycle == 180 ||
		currentCycle == 220 {
		return currentCycle * x
	}

	return 0
}

func Part1() {
	lines, err := advent.ReadInput()
	if err != nil {
		panic(err)
	}

	currentCycle := 1
	x := 1
	totalStrength := 0
	for _, line := range lines {
		if line == "noop" {
			currentCycle += 1
			totalStrength += updateStrength(x, currentCycle)
			continue
		}

		fields := strings.Split(line, " ")
		n, err := strconv.Atoi(fields[1])
		if err != nil {
			panic(err)
		}

		currentCycle += 1
		totalStrength += updateStrength(x, currentCycle)

		x += n
		currentCycle += 1
		totalStrength += updateStrength(x, currentCycle)
	}

	fmt.Println(totalStrength)
}

func main() {
	if advent.ShouldRunPart1() {
		Part1()
	}

	if advent.ShouldRunPart2() {
		Part2()
	}
}
