package main

import (
	"fmt"

	advent "github.com/BenJuan26/advent2022"
)

type stack string

func (s stack) Push(v rune) stack {
	return s + stack(v)
}

func (s stack) Pop() (stack, rune) {
	// FIXME: What do we do if the stack is empty, though?

	l := len(s)
	return s[:l-1], rune(s[l-1])
}

func (s stack) Reverse() stack {
	o := stack("")

	for i := len(s) - 1; i >= 0; i-- {
		o = o + stack(s[i])
	}

	return o
}

func Part1() {
	lines, err := advent.ReadInput()
	if err != nil {
		panic(err)
	}

	stacks := []stack{}
	for i := 0; i < 9; i++ {
		stacks = append(stacks, "")
	}

	for _, line := range lines {
		if len(line) > 0 && line[0] == '[' {
			for n := 0; n < 9; n++ {
				i := 1 + (n * 4)
				if i < len(line) && line[i] != ' ' {
					stacks[n] = stacks[n].Push(rune(line[i]))
				}
			}
		} else if len(line) > 0 && line[0:4] == "move" {

		}
	}

	for _, stack := range stacks {
		stack = stack.Reverse()
		fmt.Println(stack)
	}
	// fmt.Println("answer")
}

func main() {
	if advent.ShouldRunPart1() {
		Part1()
	}

	if advent.ShouldRunPart2() {
		Part2()
	}
}
