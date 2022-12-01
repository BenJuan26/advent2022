package main

import (
	"fmt"
	"sort"
	"strconv"

	advent "github.com/BenJuan26/advent2022"
)

func Part2() {
	lines, err := advent.ReadInput()
	if err != nil {
		panic(err)
	}

	currentElf := 0
	allElves := []int{}
	for _, line := range lines {
		if line == "" {
			allElves = append(allElves, currentElf)
			currentElf = 0
			continue
		}

		num, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		currentElf += num
	}

	sort.Ints(allElves)

	len := len(allElves)
	fmt.Println(allElves[len-1] + allElves[len-2] + allElves[len-3])
}
