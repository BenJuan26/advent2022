package main

import (
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

var ErrWrongOrder = errors.New("wrong order")

func Compare(left, right interface{}) (bool, error) {
	if l, ok := left.(int); ok {
		if r, ok := right.(int); ok {
			if l < r {
				return true, nil
			}
		} else {
			// left is int, right is list
			rList, ok := right.([]interface{})
			if !ok {
				panic(fmt.Errorf("right value %+v was not an int or a list", right))
			}
			lList := []int{l}
			return Compare(lList, rList)
		}
	} else {
		lList, ok := left.([]interface{})
		if !ok {
			panic(fmt.Errorf("left value %+v was not an int or a list", left))
		}
		if r, ok := right.(int); ok {
			if l < r {
				return Compare(lList, []int{r})
			}
		} else {
			// both lists
			rList, ok := right.([]interface{})
			if !ok {
				panic(fmt.Errorf("right value %+v was not an int or a list", right))
			}
			i := 0
			for i < len(lList) && i < len(rList) {
				result, err := Compare(lList[i], rList[i])
				if err != nil || result == true {
					return result, err
				}
				i += 1
				if i == len(lList) && i < len(rList) {
					// left ran out first
					return true, nil
				}
				if i < len(lList) && i == len(rList) {
					// right ran out first
					return false, ErrWrongOrder
				}
			}
		}
	}

	return false, errors.New("unreachable?")
}

func Part1() {
	lines, err := advent.ReadInput()
	if err != nil {
		panic(err)
	}

	total := 0
	for i := 0; i < len(lines); i += 3 {
		// remove the outside brackets since an outer list is always guaranteed
		left, _ := ParseList(lines[i])
		right, _ := ParseList(lines[i+1])

		result, err := Compare(left, right)
		if err != nil && err.Error() != ErrWrongOrder.Error() {
			panic(err)
		}
		if err == nil && result == true {
			total += (i / 3) + 1
		}
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
