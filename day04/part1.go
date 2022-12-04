package main

import (
	"fmt"
	"strconv"
	"strings"

	advent "github.com/BenJuan26/advent2022"
)

func Part1() {
	lines, err := advent.ReadInput()
	if err != nil {
		panic(err)
	}

	count := 0
	for _, line := range lines {
		rangeStrings := strings.Split(line, ",")

		range1Elements := strings.Split(rangeStrings[0], "-")
		range1Min, err := strconv.Atoi(range1Elements[0])
		if err != nil {
			panic(err)
		}
		range1Max, err := strconv.Atoi(range1Elements[1])
		if err != nil {
			panic(err)
		}

		range2Elements := strings.Split(rangeStrings[1], "-")
		range2Min, err := strconv.Atoi(range2Elements[0])
		if err != nil {
			panic(err)
		}
		range2Max, err := strconv.Atoi(range2Elements[1])
		if err != nil {
			panic(err)
		}

		if range1Min >= range2Min && range1Max <= range2Max || range2Min >= range1Min && range2Max <= range1Max {
			count += 1
		}
	}

	fmt.Println(count)
}

func main() {
	if advent.ShouldRunPart1() {
		Part1()
	}

	if advent.ShouldRunPart2() {
		Part2()
	}
}
