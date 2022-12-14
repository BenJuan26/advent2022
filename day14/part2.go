package main

import (
	"fmt"
	"strconv"
	"strings"

	advent "github.com/BenJuan26/advent2022"
)

func Part2() {
	lines, err := advent.ReadInput()
	if err != nil {
		panic(err)
	}

	maxY := 0
	visited := map[string]bool{}
	for _, line := range lines {
		points := strings.Split(line, " -> ")
		first := true
		prevPoint := Point{}
		for _, pointString := range points {
			coordStrings := strings.Split(pointString, ",")
			pointX, err := strconv.Atoi(coordStrings[0])
			if err != nil {
				panic(err)
			}
			pointY, err := strconv.Atoi(coordStrings[1])
			if err != nil {
				panic(err)
			}
			if pointY > maxY {
				maxY = pointY
			}
			point := Point{pointX, pointY}
			if first {
				prevPoint = point
				first = false
				continue
			}

			var inc func(int, int) (int, int)
			if prevPoint.X == point.X {
				if prevPoint.Y > point.Y {
					inc = func(x, y int) (int, int) {
						return x, y - 1
					}
				} else if prevPoint.Y < point.Y {
					inc = func(x, y int) (int, int) {
						return x, y + 1
					}
				}
			} else {
				if prevPoint.X > point.X {
					inc = func(x, y int) (int, int) {
						return x - 1, pointY
					}
				} else if prevPoint.X < point.X {
					inc = func(x, y int) (int, int) {
						return x + 1, y
					}
				}
			}

			for x, y := prevPoint.X, prevPoint.Y; x != point.X || y != point.Y; x, y = inc(x, y) {
				visited[Point{x, y}.String()] = true
			}
			visited[point.String()] = true
			prevPoint = point
		}
	}

	floor := maxY + 2

	done := false
	count := 0
	for !done {
		rest := false
		point := Point{500, 0}
		for !rest {
			if visited[Point{500, 0}.String()] {
				done = true
				break
			}
			if point.Y == floor-1 {
				count += 1
				visited[point.String()] = true
				rest = true
				break
			}
			p := Point{point.X, point.Y + 1}
			if !visited[p.String()] {
				point = p
				continue
			}

			p = Point{point.X - 1, point.Y + 1}
			if !visited[p.String()] {
				point = p
				continue
			}

			p = Point{point.X + 1, point.Y + 1}
			if !visited[p.String()] {
				point = p
				continue
			}

			count += 1
			visited[point.String()] = true
			rest = true
		}
	}

	fmt.Println(count)
}
