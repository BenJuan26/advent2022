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

func draw(screen []string, xReg, cycle int) {
	loopedCycle := (cycle - 1) % 240
	y := int(loopedCycle / 40)
	x := (cycle - 1) % 40
	if abs(x-xReg) > 1 {
		return
	}

	if x == 0 {
		screen[y] = "X" + screen[y][1:]
	} else if x == len(screen[y])-1 {
		screen[y] = screen[y][:x] + "X"
	} else {
		screen[y] = screen[y][:x] + "X" + screen[y][x+1:]
	}
}

func Part2() {
	lines, err := advent.ReadInput()
	if err != nil {
		panic(err)
	}

	screen := []string{}
	for i := 0; i < 6; i++ {
		screen = append(screen, strings.Repeat(".", 40))
	}

	currentCycle := 1
	x := 1
	draw(screen, x, currentCycle)
	for _, line := range lines {
		if line == "noop" {
			currentCycle += 1
			draw(screen, x, currentCycle)
			continue
		}

		fields := strings.Split(line, " ")
		n, err := strconv.Atoi(fields[1])
		if err != nil {
			panic(err)
		}

		currentCycle += 1
		draw(screen, x, currentCycle)

		x += n
		currentCycle += 1
		draw(screen, x, currentCycle)
	}

	for _, line := range screen {
		fmt.Println(line)
	}
}
