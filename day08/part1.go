package main

import (
	"fmt"
	"strconv"

	advent "github.com/BenJuan26/advent2022"
)

type Tree struct {
	Height         int
	AlreadyVisible bool
}

func Part1() {
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

	totalVisible := 0
	// from left
	for col := 0; col < len(grid); col++ {
		highest := -1
		for row := 0; row < len(grid[0]); row++ {
			if tree := grid[col][row]; tree.Height > highest {
				if !tree.AlreadyVisible {
					totalVisible += 1
				}
				tree.AlreadyVisible = true
				highest = tree.Height
				if tree.Height == 9 {
					break
				}
			}
		}
	}

	// from top
	for row := 0; row < len(grid[0]); row++ {
		highest := -1
		for col := 0; col < len(grid); col++ {
			if tree := grid[col][row]; tree.Height > highest {
				if !tree.AlreadyVisible {
					totalVisible += 1
				}
				tree.AlreadyVisible = true
				highest = tree.Height
				if tree.Height == 9 {
					break
				}
			}
		}
	}

	// from right
	for col := len(grid) - 1; col >= 0; col-- {
		highest := -1
		for row := len(grid[0]) - 1; row >= 0; row-- {
			if tree := grid[col][row]; tree.Height > highest {
				if !tree.AlreadyVisible {
					totalVisible += 1
				}
				tree.AlreadyVisible = true
				highest = tree.Height
				if tree.Height == 9 {
					break
				}
			}
		}
	}

	// from bottom
	for row := len(grid[0]) - 1; row >= 0; row-- {
		highest := -1
		for col := len(grid) - 1; col >= 0; col-- {
			if tree := grid[col][row]; tree.Height > highest {
				if !tree.AlreadyVisible {
					totalVisible += 1
				}
				tree.AlreadyVisible = true
				highest = tree.Height
				if tree.Height == 9 {
					break
				}
			}
		}
	}

	fmt.Println(totalVisible)
}

func main() {
	if advent.ShouldRunPart1() {
		Part1()
	}

	if advent.ShouldRunPart2() {
		Part2()
	}
}
