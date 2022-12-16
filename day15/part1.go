package main

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"

	advent "github.com/BenJuan26/advent2022"
)

type Point struct {
	X int
	Y int
}

type Sensor struct {
	Position          Point
	NearestBeacon     Point
	ManhattanDistance int
}

func Part1() {
	lines, err := advent.ReadInput()
	if err != nil {
		panic(err)
	}

	lineRegex := regexp.MustCompile(`Sensor at x=(\-?)(\d+), y=(\-?)(\d+): closest beacon is at x=(\-?)(\d+), y=(\-?)(\d+)`)

	targetRow := 2000000
	sensors := []Sensor{}
	unavailable := map[int]bool{}
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
		if advent.Abs(targetRow-s.Position.Y) <= s.ManhattanDistance {
			yDiff := advent.Abs(s.Position.Y - targetRow)
			width := 2*(s.ManhattanDistance-yDiff) + 1
			for i := s.Position.X - ((width - 1) / 2); i <= s.Position.X+((width-1)/2); i++ {
				unavailable[i] = true
			}
		}
	}

	for _, sensor := range sensors {
		if sensor.NearestBeacon.Y == targetRow {
			delete(unavailable, sensor.NearestBeacon.X)
		}
	}

	fmt.Println(len(unavailable))
}

func main() {
	if advent.ShouldRunPart1() {
		Part1()
	}

	if advent.ShouldRunPart2() {
		Part2()
	}
}
