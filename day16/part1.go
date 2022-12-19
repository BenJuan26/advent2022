package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"

	advent "github.com/BenJuan26/advent2022"
)

type Valve struct {
	Name     string
	FlowRate int
	Tunnels  []string
	Score    int
}

func MinDist(q map[string]Valve, dist map[string]int) Valve {
	min := math.MaxInt
	var best Valve
	// just set best to the first valve in the set
	for _, v := range q {
		best = v
		break
	}

	for _, v := range q {
		d := dist[v.Name]
		if d < min {
			min = d
			best = v
		}
	}

	return best
}

func ReconstructPath(prev map[string]Valve, current Valve) []Valve {
	totalPath := []Valve{current}
	_, ok := prev[current.Name]
	for ok {
		current = prev[current.Name]
		totalPath = append([]Valve{current}, totalPath...)
		_, ok = prev[current.Name]
	}

	return totalPath
}

func Dijkstra(valves map[string]Valve, source Valve) (map[string]int, map[string]Valve) {
	dist := map[string]int{}
	prev := map[string]Valve{}
	q := map[string]Valve{}
	for _, v := range valves {
		dist[v.Name] = math.MaxInt
		q[v.Name] = v
	}

	dist[source.Name] = 0

	for len(q) > 0 {
		u := MinDist(q, dist)
		delete(q, u.Name)
		for _, vName := range u.Tunnels {
			v, ok := q[vName]
			if !ok {
				continue
			}

			alt := dist[u.Name] + 1
			if alt < dist[v.Name] {
				dist[v.Name] = alt
				prev[v.Name] = u
			}
		}
	}

	return dist, prev
}

func TryOptions(valves map[string]Valve, source string, timeElapsed, totalPressure int, opened_ map[string]bool) (int, int) {
	if timeElapsed > 30 {
		return timeElapsed, totalPressure
	}

	dist, _ := Dijkstra(valves, valves[source])
	scores := []Valve{}
	for vName, distToValve := range dist {
		if opened_[vName] {
			continue
		}
		v := valves[vName]
		v.Score = (30 - distToValve) * v.FlowRate
		scores = append(scores, v)
	}

	sort.Slice(scores, func(i, j int) bool {
		return scores[i].Score > scores[j].Score
	})

	// try the 12 best options
	maxTime := timeElapsed
	maxPressure := math.MinInt
	found := false
	for i, v := range scores {
		if i > 11 {
			break
		}

		opened := map[string]bool{}
		for k, v := range opened_ {
			opened[k] = v
		}
		opened[v.Name] = true

		t := timeElapsed + dist[v.Name] + 1
		time, pressure := TryOptions(valves, v.Name, t,
			totalPressure+(30-t)*v.FlowRate, opened)
		if pressure > maxPressure && time <= 30 {
			maxTime = time
			maxPressure = pressure
			found = true
		}
	}

	if !found {
		return timeElapsed, totalPressure
	}

	return maxTime, maxPressure
}

func Part1() {
	lines, err := advent.ReadInput()
	if err != nil {
		panic(err)
	}

	valves := map[string]Valve{}

	for _, line := range lines {
		fields := strings.Split(line, " ")
		valveName := fields[1]
		rateString := fields[4]
		rateString = rateString[5 : len(rateString)-1]
		rate, err := strconv.Atoi(rateString)
		if err != nil {
			panic(err)
		}

		tunnelStrings := fields[9:]
		tunnels := []string{}
		for _, tunnel := range tunnelStrings {
			replacer := strings.NewReplacer(",", "")
			tunnel = replacer.Replace(tunnel)
			tunnels = append(tunnels, tunnel)
		}

		valves[valveName] = Valve{valveName, rate, tunnels, 0}
	}

	opened := map[string]bool{"AA": true}
	_, totalPressure := TryOptions(valves, "AA", 0, 0, opened)

	fmt.Println(totalPressure)
}

func main() {
	if advent.ShouldRunPart1() {
		Part1()
	}

	if advent.ShouldRunPart2() {
		Part2()
	}
}
