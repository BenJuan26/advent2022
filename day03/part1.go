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

	total := 0
	for _, line := range lines {
		compartment1 := line[:len(line)/2]
		compartment2 := line[len(line)/2:]

		contents := map[rune]bool{}
		for _, item := range compartment1 {
			contents[item] = true
		}

		var commonItem rune
		for _, item := range compartment2 {
			if _, ok := contents[item]; ok {
				commonItem = item
				break
			}
		}

		value := 0
		if commonItem < 'a' {
			value = int(commonItem - 'A' + 27)
		} else {
			value = int(commonItem - 'a' + 1)
		}

		total += value
	}

	fmt.Println(total)
}

func main() {
	if advent.ShouldRunPart1() {
		Part1()
	}

	if advent.ShouldRunPart2() {
		Part2()
	}
}
