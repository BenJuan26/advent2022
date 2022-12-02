package main

import (
	"fmt"
	"strings"

	advent "github.com/BenJuan26/advent2022"
)

func Part1() {
	lines, err := advent.ReadInput()
	if err != nil {
		panic(err)
	}

	wins := map[string]string{
		"A": "Y",
		"B": "Z",
		"C": "X",
	}

	ties := map[string]string{
		"A": "X",
		"B": "Y",
		"C": "Z",
	}

	moveScores := map[string]int{
		"X": 1,
		"Y": 2,
		"Z": 3,
	}

	totalScore := 0
	for _, line := range lines {
		moves := strings.Split(line, " ")
		theirMove := moves[0]
		myMove := moves[1]

		outcomeScore := 0

		if wins[theirMove] == myMove {
			outcomeScore = 6
		} else if ties[theirMove] == myMove {
			outcomeScore = 3
		}

		moveScore := moveScores[myMove]
		totalScore += outcomeScore + moveScore
	}

	fmt.Println(totalScore)
}

func main() {
	if advent.ShouldRunPart1() {
		Part1()
	}

	if advent.ShouldRunPart2() {
		Part2()
	}
}
