package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	advent "github.com/BenJuan26/advent2022"
)

func Part2() {
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

		maxOfMins := max(range1Min, range2Min)
		minOfMaxes := min(range1Max, range2Max)

		if maxOfMins <= minOfMaxes {
			count += 1
		}
	}

	fmt.Println(count)
}

func max(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}

func min(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}
