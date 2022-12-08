package main

import (
	"fmt"
	"strconv"

	advent "github.com/BenJuan26/advent2022"
)

func getScore(grid [][]*Tree, c int, r int) int {
	up := 0
	for col := c - 1; col >= 0; col-- {
		if grid[col][r].Height < grid[c][r].Height {
			up += 1
		} else {
			up += 1
			break
		}
	}

	down := 0
	for col := c + 1; col < len(grid); col++ {
		if grid[col][r].Height < grid[c][r].Height {
			down += 1
		} else {
			down += 1
			break
		}
	}

	left := 0
	for row := r - 1; row >= 0; row-- {
		if grid[c][row].Height < grid[c][r].Height {
			left += 1
		} else {
			left += 1
			break
		}
	}

	right := 0
	for row := r + 1; row < len(grid[0]); row++ {
		if grid[c][row].Height < grid[c][r].Height {
			right += 1
		} else {
			right += 1
			break
		}
	}

	return up * down * left * right
}

func Part2() {
	lines, err := advent.ReadInput()
	if err != nil {
		panic(err)
	}

	grid := [][]*Tree{}
	for _, line := range lines {
		row := []*Tree{}
		for _, char := range line {
			num, err := strconv.Atoi(string(char))
			if err != nil {
				panic(err)
			}

			row = append(row, &Tree{num, false})
		}
		grid = append(grid, row)
	}

	bestScore := 0
	for col := 0; col < len(grid); col++ {
		for row := 0; row < len(grid[0]); row++ {
			score := getScore(grid, col, row)
			if score > bestScore {
				bestScore = score
			}
		}
	}

	fmt.Println(bestScore)
}
