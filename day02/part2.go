package main

import (
	"fmt"
	"strings"

	advent "github.com/BenJuan26/advent2022"
)

func Part2() {
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

	losses := map[string]string{
		"A": "Z",
		"B": "X",
		"C": "Y",
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
		outcome := moves[1]

		outcomeScore := 0
		myMove := ""

		if outcome == "X" {
			myMove = losses[theirMove]
		} else if outcome == "Y" {
			myMove = ties[theirMove]
			outcomeScore = 3
		} else if outcome == "Z" {
			myMove = wins[theirMove]
			outcomeScore = 6
		}

		moveScore := moveScores[myMove]
		totalScore += outcomeScore + moveScore
	}

	fmt.Println(totalScore)
}
