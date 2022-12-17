package main

import (
	"errors"
	"fmt"
	"regexp"
	"sort"
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

type Range struct {
	Min int
	Max int
}

func (r Range) Contains(x int) bool {
	return x >= r.Min && x <= r.Max
}

func CanReduce(a, b Range) bool {
	return a.Contains(b.Min) || a.Contains(b.Max) || b.Contains(a.Min) || b.Contains(a.Max) ||
		b.Min == a.Max+1 || a.Min == b.Max+1
}

func Reduce(a, b Range) Range {
	return Range{advent.Min(a.Min, b.Min), advent.Max(a.Max, b.Max)}
}

type Ranges []Range

func (r Ranges) Len() int           { return len(r) }
func (r Ranges) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
func (r Ranges) Less(i, j int) bool { return r[i].Min < r[j].Min }

func (r Ranges) ReduceAll() Ranges {
	limit := len(r)
	for i := 0; i < limit; i++ {
		for j := i + 1; j < limit; j++ {
			a := r[i]
			b := r[j]
			if CanReduce(a, b) {
				r[i] = Reduce(a, b)
				r = append(r[:j], r[j+1:]...)
				j -= 1
				limit -= 1
			}
		}
	}

	return r
}

func (r Ranges) Push(a Range) Ranges {
	r = append(r, a)
	sort.Sort(r)
	return r
}

// Total returns the total integers contained in the ranges.
// The ranges must be sorted and reduced.
func (r Ranges) Total() int {
	total := 0
	for _, a := range r {
		total += a.Max - a.Min
	}

	return total
}

func NewRanges() Ranges {
	return Ranges([]Range{})
}

func Part1() {
	lines, err := advent.ReadInput()
	if err != nil {
		panic(err)
	}

	lineRegex := regexp.MustCompile(`Sensor at x=(\-?)(\d+), y=(\-?)(\d+): closest beacon is at x=(\-?)(\d+), y=(\-?)(\d+)`)

	targetRow := 2000000
	ranges := NewRanges()
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
		if advent.Abs(targetRow-s.Position.Y) <= s.ManhattanDistance {
			yDiff := advent.Abs(s.Position.Y - targetRow)
			width := 2*(s.ManhattanDistance-yDiff) + 1
			minX := s.Position.X - ((width - 1) / 2)
			maxX := s.Position.X + ((width - 1) / 2)
			ranges = ranges.Push(Range{minX, maxX})
			ranges = ranges.ReduceAll()
		}
	}

	fmt.Println(ranges.Total())
}

func main() {
	if advent.ShouldRunPart1() {
		Part1()
	}

	if advent.ShouldRunPart2() {
		Part2()
	}
}
