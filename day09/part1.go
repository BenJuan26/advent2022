package main

import (
	"fmt"
	"strconv"
	"strings"

	advent "github.com/BenJuan26/advent2022"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

type Knot struct {
	X int
	Y int
}

func (k *Knot) Move(direction string) {
	switch direction {
	case "R":
		k.X += 1
	case "L":
		k.X -= 1
	case "U":
		k.Y -= 1
	case "D":
		k.Y += 1
	}
}

func (k *Knot) Follow(head *Knot) {
	if k.X < head.X {
		k.X += 1
	} else if k.X > head.X {
		k.X -= 1
	}

	if k.Y < head.Y {
		k.Y += 1
	} else if k.Y > head.Y {
		k.Y -= 1
	}
}

func SimulateRope(ropeLength int) int {
	lines, err := advent.ReadInput()
	if err != nil {
		panic(err)
	}

	knots := []*Knot{}
	for i := 0; i < ropeLength; i++ {
		knots = append(knots, &Knot{})
	}
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

		for step := 0; step < distance; step++ {
			// move the head
			knots[0].Move(direction)
			for i := 1; i < ropeLength; i++ {
				head := knots[i-1] // "head" in this case means the knot ahead of the current one
				tail := knots[i]
				dx := head.X - tail.X
				dy := head.Y - tail.Y
				if abs(dx) > 1 || abs(dy) > 1 {
					tail.Follow(head)
					if _, ok := visited[fmt.Sprintf("%d,%d", tail.X, tail.Y)]; !ok && i == len(knots)-1 {
						visited[fmt.Sprintf("%d,%d", tail.X, tail.Y)] = true
						numVisited += 1
					}
				}
			}
		}
	}

	return numVisited
}

func Part1() {
	fmt.Println(SimulateRope(2))
}

func main() {
	if advent.ShouldRunPart1() {
		Part1()
	}

	if advent.ShouldRunPart2() {
		Part2()
	}
}
