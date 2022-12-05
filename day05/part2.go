package main

import (
	"fmt"
	"regexp"
	"strconv"

	advent "github.com/BenJuan26/advent2022"
)

func Part2() {
	lines, err := advent.ReadInput()
	if err != nil {
		panic(err)
	}

	stacks := []stack{}
	for i := 0; i < 9; i++ {
		stacks = append(stacks, "")
	}

	re := regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`)
	reversed := false
	for _, line := range lines {
		if len(line) > 0 && line[0] == '[' {
			for n := 0; n < 9; n++ {
				i := 1 + (n * 4)
				if i < len(line) && line[i] != ' ' {
					stacks[n] = stacks[n].Push(rune(line[i]))
				}
			}
		} else if len(line) > 0 && line[0:4] == "move" {
			match := re.FindStringSubmatch(line)
			amount, err := strconv.Atoi(match[1])
			if err != nil {
				panic(err)
			}

			from, err := strconv.Atoi(match[2])
			if err != nil {
				panic(err)
			}

			to, err := strconv.Atoi(match[3])
			if err != nil {
				panic(err)
			}

			from -= 1
			to -= 1

			tempStack := stack("")
			for i := 0; i < amount; i++ {
				var value rune
				stacks[from], value = stacks[from].Pop()
				tempStack = tempStack.Push(value)
			}
			tempStack = tempStack.Reverse()
			stacks[to] = stacks[to] + tempStack
		} else if !reversed {
			for i := range stacks {
				stacks[i] = stacks[i].Reverse()
			}
			reversed = true
		}
	}

	for _, stack := range stacks {
		fmt.Print(string(stack[len(stack)-1]))
	}
	fmt.Printf("\n")
}
