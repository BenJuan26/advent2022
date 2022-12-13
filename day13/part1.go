package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	advent "github.com/BenJuan26/advent2022"
)

func ParseList(s string) ([]interface{}, int) {
	var items []interface{}
	// start at 1: we know the first is '['
	tempNum := ""
	i := 1
	for i < len(s) {
		if s[i] == '[' {
			parsed, size := ParseList(s[i:])
			items = append(items, parsed)
			i += size
		} else if s[i] == ',' {
			if tempNum != "" {
				num, err := strconv.Atoi(tempNum)
				if err != nil {
					panic(err)
				}
				items = append(items, num)
				tempNum = ""
			}
		} else if s[i] == ']' {
			if tempNum != "" {
				num, err := strconv.Atoi(tempNum)
				if err != nil {
					panic(err)
				}
				items = append(items, num)
				tempNum = ""
			}
			return items, i
		} else {
			// this should always be a digit
			tempNum += string(s[i])
		}
		i += 1
	}

	panic(errors.New("should have found a closing bracket"))
}

func Part1() {
	lines, err := advent.ReadInput()
	if err != nil {
		panic(err)
	}

	// for i := 0; i < len(lines); i += 3 {
	// 	// remove the outside brackets since an outer list is always guaranteed
	// 	left := ParseList(lines[i])
	// 	right := lines[1 : len(lines[i+1])-2]

	// }

	parsed, _ := ParseList(lines[21])
	b, err := json.Marshal(parsed)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
}

func main() {
	if advent.ShouldRunPart1() {
		Part1()
	}

	if advent.ShouldRunPart2() {
		Part2()
	}
}
