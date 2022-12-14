package main

import (
	"fmt"
	"math"

	advent "github.com/BenJuan26/advent2022"
)

type Point struct {
	X      int
	Y      int
	Height rune
	FScore float64
}

func (p *Point) String() string {
	return fmt.Sprintf("%d,%d", p.X, p.Y)
}

func Part1() {
	lines, err := advent.ReadInput()
	if err != nil {
		panic(err)
	}

	var start, goal *Point
	for y, line := range lines {
		for x, char := range line {
			if char == 'S' {
				start = &Point{x, y, 'a', math.Inf(1)}
				lines[y] = line[:x] + "a" + line[x+1:]
			}

			if char == 'E' {
				goal = &Point{x, y, 'z', math.Inf(1)}
				lines[y] = line[:x] + "z" + line[x+1:]
			}
		}
	}

	path, err := AStar(start, goal, TaxiCab, lines)
	if err != nil {
		panic(err)
	}

	fmt.Println(len(path) - 1)
}

func main() {
	if advent.ShouldRunPart1() {
		Part1()
	}

	if advent.ShouldRunPart2() {
		Part2()
	}
}
