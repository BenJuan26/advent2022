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

	total := 0
	i := 0
	for i < len(lines) {
		contents := map[rune]bool{}
		for _, item := range lines[i] {
			contents[item] = true
		}

		commonContents := map[rune]bool{}
		for _, item := range lines[i+1] {
			if _, ok := contents[item]; ok {
				commonContents[item] = true
			}
		}

		var commonItem rune
		for _, item := range lines[i+2] {
			if _, ok := commonContents[item]; ok {
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
		i += 3
	}

	fmt.Println(total)
}
