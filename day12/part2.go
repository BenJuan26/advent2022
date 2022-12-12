package main

import (
	"fmt"

	advent "github.com/BenJuan26/advent2022"
)

func Part2() {
	lines, err := advent.ReadInput()
	if err != nil {
		panic(err)
	}

	fmt.Println(lines[0])
}
