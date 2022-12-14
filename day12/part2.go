package main

import (
	"fmt"
	"math"

	advent "github.com/BenJuan26/advent2022"
)

func Part2() {
	lines, err := advent.ReadInput()
	if err != nil {
		panic(err)
	}

	var start, goal *Point
	starts := []*Point{}
	minDistance := math.MaxInt
	for y, line := range lines {
		for x, char := range line {
			if char == 'S' {
				start = &Point{x, y, 'a', math.Inf(1)}
				lines[y] = line[:x] + "a" + line[x+1:]
				starts = append(starts, start)
			}

			if char == 'E' {
				goal = &Point{x, y, 'z', math.Inf(1)}
				lines[y] = line[:x] + "z" + line[x+1:]
			}

			if char == 'a' {
				starts = append(starts, &Point{x, y, char, math.Inf(1)})
			}
		}
	}

	for _, s := range starts {
		path, err := AStar(s, goal, TaxiCab, lines)
		if err != nil {
			// goal is unreachable from this position
			continue
		}

		if len(path)-1 < minDistance {
			minDistance = len(path) - 1
		}
	}

	fmt.Println(minDistance)
}
