package main

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"

	advent "github.com/BenJuan26/advent2022"
)

func Part2() {
	lines, err := advent.ReadInput()
	if err != nil {
		panic(err)
	}

	lineRegex := regexp.MustCompile(`Sensor at x=(\-?)(\d+), y=(\-?)(\d+): closest beacon is at x=(\-?)(\d+), y=(\-?)(\d+)`)

	sensors := []Sensor{}
	for _, line := range lines {
		match := lineRegex.FindStringSubmatch(line)
		if match == nil {
			panic(fmt.Errorf("regex didn't match for line %s", line))
		}

		if len(match) < 9 {
			panic(errors.New("not enough matches"))
		}

		values := []int{}
		for i := 1; i <= 8; i += 2 {
			stringValue := match[i+1]
			value, err := strconv.Atoi(stringValue)
			if err != nil {
				panic(err)
			}
			if match[i] == "-" {
				value = -value
			}
			values = append(values, value)
		}

		manhattan := advent.Abs(values[0]-values[2]) + advent.Abs(values[1]-values[3])

		s := Sensor{Point{values[0], values[1]}, Point{values[2], values[3]}, manhattan}
		sensors = append(sensors, s)
	}

	maxCoord := 4000000
	var x, y int
	found := false
	for y = 0; y <= maxCoord; y++ {
		ranges := NewRanges()
		for _, s := range sensors {
			if advent.Abs(y-s.Position.Y) <= s.ManhattanDistance {
				yDiff := advent.Abs(s.Position.Y - y)
				width := 2*(s.ManhattanDistance-yDiff) + 1
				minX := advent.Max(s.Position.X-((width-1)/2), 0)
				maxX := advent.Min(s.Position.X+((width-1)/2), maxCoord)
				ranges = ranges.Push(Range{minX, maxX})
				ranges = ranges.ReduceAll()
			}
		}
		// This should only be the case when there is a gap in the ranges
		// otherwise they should all be reduced to a single range
		if len(ranges) == 2 {
			x = ranges[0].Max + 1
			found = true
			break
		}
	}

	if !found {
		panic("not found")
	}

	fmt.Println(x*4000000 + y)
}
