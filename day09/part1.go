package main

import (
	"fmt"
	"strconv"
	"strings"

	advent "github.com/BenJuan26/advent2022"
)

func moveHead(hx, hy int, direction string) (int, int) {
	if direction == "R" {
		return hx - 1, hy
	}

	if direction == "L" {
		return hx + 1, hy
	}

	if direction == "U" {
		return hx, hy - 1
	}

	return hx, hy + 1
}

func moveTail(hx, hy int, direction string) (int, int) {
	if direction == "R" {
		return hx + 1, hy
	}

	if direction == "L" {
		return hx - 1, hy
	}

	if direction == "U" {
		return hx, hy + 1
	}

	return hx, hy - 1
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func Part1() {
	lines, err := advent.ReadInput()
	if err != nil {
		panic(err)
	}

	var hx, hy, tx, ty int
	visited := map[string]bool{}
	numVisited := 1 // the initial position
	for _, line := range lines {
		fields := strings.Split(line, " ")
		direction := fields[0]
		distanceString := fields[1]
		distance, err := strconv.Atoi(distanceString)
		if err != nil {
			panic(err)
		}

		for i := 0; i < distance; i++ {
			hx, hy = moveHead(hx, hy, direction)
			dx := hx - tx
			dy := hy - ty
			if abs(dx) > 1 || abs(dy) > 1 {
				tx, ty = moveTail(hx, hy, direction)
				// fmt.Printf("%d, %d\n", tx, ty)
				if _, ok := visited[fmt.Sprintf("%d,%d", tx, ty)]; !ok {
					visited[fmt.Sprintf("%d,%d", tx, ty)] = true
					numVisited += 1
				}
			}
		}
	}

	fmt.Println(numVisited)
}

func main() {
	if advent.ShouldRunPart1() {
		Part1()
	}

	if advent.ShouldRunPart2() {
		Part2()
	}
}
