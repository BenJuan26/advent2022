package main

import (
	"errors"
	"math"
)

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

func TaxiCab(a, b *Point) float64 {
	return float64(abs(a.X-b.X) + abs(a.Y-b.Y))
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
		n = append(n, &Point{x, y, rune(lines[y][x]), math.Inf(1)})
	}

	if p.X < len(lines[0])-1 {
		x := p.X + 1
		y := p.Y
		n = append(n, &Point{x, y, rune(lines[y][x]), math.Inf(1)})
	}

	if p.Y > 0 {
		x := p.X
		y := p.Y - 1
		n = append(n, &Point{x, y, rune(lines[y][x]), math.Inf(1)})
	}

	if p.Y < len(lines)-1 {
		x := p.X
		y := p.Y + 1
		n = append(n, &Point{x, y, rune(lines[y][x]), math.Inf(1)})
	}

	return n
}

func Distance(a, b *Point) float64 {
	if b.Height <= a.Height+1 {
		return 1
	}

	return math.Inf(1)
}

type Heuristic func(start, goal *Point) float64

func AStar(start, goal *Point, h Heuristic, lines []string) ([]*Point, error) {
	openSet := NewQ()
	start.FScore = float64(h(start, goal))
	openSet.Push(start)

	cameFrom := map[string]*Point{}

	gScore := map[string]float64{}
	gScore[start.String()] = 0

	for !openSet.IsEmpty() {
		current := openSet.Pop()
		if current.String() == goal.String() {
			return ReconstructPath(cameFrom, current), nil
		}

		_, hasGScore := gScore[current.String()]
		if !hasGScore {
			gScore[current.String()] = math.Inf(1)
		}
		neigbours := GetNeighbours(current, lines)
		for _, n := range neigbours {
			_, hasGScore = gScore[n.String()]
			if !hasGScore {
				gScore[n.String()] = math.Inf(1)
			}
			tentativeGScore := gScore[current.String()] + Distance(current, n)
			if tentativeGScore < gScore[n.String()] {
				cameFrom[n.String()] = current
				gScore[n.String()] = tentativeGScore
				n.FScore = tentativeGScore + h(n, goal)
				if !openSet.Contains(n) {
					openSet.Push(n)
				}
			}
		}
	}

	return nil, errors.New("openSet is empty but goal was never reached")
}
