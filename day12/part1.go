package main

import (
	"errors"
	"fmt"
	"math"

	advent "github.com/BenJuan26/advent2022"
)

type Point struct {
	X      int
	Y      int
	Height rune
	FScore int
}

func (p *Point) String() string {
	return fmt.Sprintf("%d,%d", p.X, p.Y)
}

type NodeQueue struct {
	Items []*Point
	Map   map[string]bool
}

func (s *NodeQueue) Push(p *Point) {
	s.Map[p.String()] = true
	if len(s.Items) == 0 {
		s.Items = append(s.Items, p)
		return
	}

	inserted := false
	for i, item := range s.Items {
		if p.FScore < item.FScore {
			inserted = true
			if i > 0 {
				s.Items = append(s.Items[:i+1], s.Items[i:]...)
				s.Items[i] = p
			} else {
				s.Items = append([]*Point{p}, s.Items...)
			}
		}
	}

	if !inserted {
		s.Items = append(s.Items, p)
	}
}

func (s *NodeQueue) Pop() *Point {
	item := s.Items[0]
	s.Items = s.Items[1:]
	delete(s.Map, item.String())
	return item
}

func (s *NodeQueue) Contains(p *Point) bool {
	_, ok := s.Map[p.String()]
	return ok
}

func (s *NodeQueue) IsEmpty() bool {
	return len(s.Items) == 0
}

func (s *NodeQueue) Size() int {
	return len(s.Items)
}

func NewQ() *NodeQueue {
	s := &NodeQueue{}
	s.Items = []*Point{}
	s.Map = map[string]bool{}
	return s
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func TaxiCab(a, b *Point) int {
	return abs(a.X-b.X) + abs(a.Y-b.Y)
}

func ReconstructPath(cameFrom map[string]*Point, current *Point) []*Point {
	totalPath := []*Point{current}
	_, ok := cameFrom[current.String()]
	for ok {
		current = cameFrom[current.String()]
		totalPath = append([]*Point{current}, totalPath...)
		_, ok = cameFrom[current.String()]
	}

	return totalPath
}

func GetNeighbours(p *Point, lines []string) []*Point {
	n := []*Point{}
	if p.X > 0 {
		x := p.X - 1
		y := p.Y
		n = append(n, &Point{x, y, rune(lines[y][x]), math.MaxInt})
	}

	if p.X < len(lines[0])-1 {
		x := p.X + 1
		y := p.Y
		n = append(n, &Point{x, y, rune(lines[y][x]), math.MaxInt})
	}

	if p.Y > 0 {
		x := p.X
		y := p.Y - 1
		n = append(n, &Point{x, y, rune(lines[y][x]), math.MaxInt})
	}

	if p.Y < len(lines)-1 {
		x := p.X
		y := p.Y + 1
		n = append(n, &Point{x, y, rune(lines[y][x]), math.MaxInt})
	}

	return n
}

func Distance(a, b *Point) int {
	if b.Height <= a.Height-1 {
		return 1
	}

	return math.MaxInt
}

type Heuristic func(start, goal *Point) int

func AStar(start, goal *Point, h Heuristic, lines []string) ([]*Point, error) {
	openSet := NewQ()
	openSet.Push(start)

	cameFrom := map[string]*Point{}

	gScore := map[string]int{}
	gScore[start.String()] = 0

	fScore := map[string]int{}
	fScore[start.String()] = TaxiCab(start, goal)

	for !openSet.IsEmpty() {
		current := openSet.Pop()
		if current.String() == goal.String() {
			return ReconstructPath(cameFrom, current), nil
		}

		neigbours := GetNeighbours(current, lines)
		for _, n := range neigbours {
			tentativeGScore := gScore[current.String()] + Distance(current, n)
			if tentativeGScore < gScore[n.String()] {
				cameFrom[n.String()] = current
				gScore[n.String()] = tentativeGScore
				fScore[n.String()] = tentativeGScore + h(n, goal)
				if !openSet.Contains(n) {
					openSet.Push(n)
				}
			}
		}
	}

	return nil, errors.New("openSet is empty but goal was never reached")
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
				start = &Point{x, y, 'a', math.MaxInt}
			}

			if char == 'E' {
				goal = &Point{x, y, 'z', math.MaxInt}
			}
		}
	}

	path, err := AStar(start, goal, TaxiCab, lines)
	if err != nil {
		panic(err)
	}

	fmt.Println(len(path))
}

func main() {
	if advent.ShouldRunPart1() {
		Part1()
	}

	if advent.ShouldRunPart2() {
		Part2()
	}
}
