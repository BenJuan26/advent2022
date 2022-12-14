package main

import (
	"errors"
	"fmt"
	"strconv"

	advent "github.com/BenJuan26/advent2022"
)

type ListItem struct {
	Value  int
	List   []ListItem
	IsList bool
}

func (l ListItem) Len() int {
	return len(l.List)
}

func (l ListItem) Less(i, j int) bool {
	result, err := Compare(l.List[i], l.List[j])
	return result && err == nil
}

func (l ListItem) Swap(i, j int) {
	l.List[i], l.List[j] = l.List[j], l.List[i]
}

func NewList() ListItem {
	l := ListItem{}
	l.List = []ListItem{}
	l.IsList = true

	return l
}

func WrapWithList(v ListItem) ListItem {
	return ListItem{0, []ListItem{v}, true}
}

func ParseList(s string) (ListItem, int) {
	l := NewList()
	// start at 1: we know the first is '['
	tempNum := ""
	i := 1
	for i < len(s) {
		if s[i] == '[' {
			parsed, size := ParseList(s[i:])
			l.List = append(l.List, parsed)
			i += size
		} else if s[i] == ',' {
			if tempNum != "" {
				num, err := strconv.Atoi(tempNum)
				if err != nil {
					panic(err)
				}
				l.List = append(l.List, ListItem{num, nil, false})
				tempNum = ""
			}
		} else if s[i] == ']' {
			if tempNum != "" {
				num, err := strconv.Atoi(tempNum)
				if err != nil {
					panic(err)
				}
				l.List = append(l.List, ListItem{num, nil, false})
				tempNum = ""
			}
			return l, i
		} else {
			// this should always be a digit
			tempNum += string(s[i])
		}
		i += 1
	}

	panic(errors.New("should have found a closing bracket"))
}

var ErrWrongOrder = errors.New("wrong order")

func Compare(left, right ListItem) (bool, error) {
	if !left.IsList {
		if !right.IsList {
			// both ints
			if left.Value < right.Value {
				// we're done, the order is correct
				return true, nil
			}

			if left.Value > right.Value {
				// we're done, the order is not correct
				return false, ErrWrongOrder
			}

			// the order is not incorrect yet, but we're not done
			return false, nil
		} else {
			// left is int, right is list
			return Compare(WrapWithList(left), right)
		}
	} else {
		if !right.IsList {
			// left is list, right is int
			return Compare(left, WrapWithList(right))
		} else {
			// both lists
			if len(left.List) == 0 && len(right.List) > 0 {
				// left ran out first
				return true, nil
			}
			if len(left.List) > 0 && len(right.List) == 0 {
				// right ran out first
				return false, ErrWrongOrder
			}
			i := 0
			for i < len(left.List) && i < len(right.List) {
				result, err := Compare(left.List[i], right.List[i])
				if err != nil || result {
					return result, err
				}
				i += 1
				if i == len(left.List) && i < len(right.List) {
					// left ran out first
					return true, nil
				}
				if i < len(left.List) && i == len(right.List) {
					// right ran out first
					return false, ErrWrongOrder
				}
			}
		}
	}

	return false, nil
}

func Part1() {
	lines, err := advent.ReadInput()
	if err != nil {
		panic(err)
	}

	total := 0
	for i := 0; i < len(lines); i += 3 {
		left, _ := ParseList(lines[i])
		right, _ := ParseList(lines[i+1])

		result, err := Compare(left, right)
		if err != nil && !errors.Is(err, ErrWrongOrder) {
			panic(err)
		}

		if err == nil && result == true {
			pairIndex := (i / 3) + 1
			total += pairIndex
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
