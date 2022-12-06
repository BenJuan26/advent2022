package main

import (
	"fmt"

	advent "github.com/BenJuan26/advent2022"
)

func Part1() {
	lines, err := advent.ReadInput()
	if err != nil {
		panic(err)
	}

	line := lines[0]
	index := 0

	for i := 4; i < len(line); i++ {
		buffer := line[i-4 : i]

		used := map[rune]bool{}
		marker := true
		for _, char := range buffer {
			if _, ok := used[char]; ok {
				marker = false
				break
			}
			used[char] = true
		}
		if marker {
			index = i
			break
		}
	}

	fmt.Println(index)
}

func main() {
	if advent.ShouldRunPart1() {
		Part1()
	}

	if advent.ShouldRunPart2() {
		Part2()
	}
}
