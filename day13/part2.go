package main

import (
	"errors"
	"fmt"
	"sort"

	advent "github.com/BenJuan26/advent2022"
)

func Part2() {
	lines, err := advent.ReadInput()
	if err != nil {
		panic(err)
	}

	bigList := ListItem{0, []ListItem{}, true}
	for _, line := range lines {
		if line == "" {
			continue
		}
		// remove the outside brackets since an outer list is always guaranteed
		packet, _ := ParseList(line)
		bigList.List = append(bigList.List, packet)
	}

	// add divider packets
	bigList.List = append(bigList.List, ListItem{0, []ListItem{{2, nil, false}}, true})
	bigList.List = append(bigList.List, ListItem{0, []ListItem{{6, nil, false}}, true})

	sort.Sort(bigList)

	firstIndex := -1
	secondIndex := -1
	for i, packet := range bigList.List {
		if len(packet.List) == 1 {
			if packet.List[0].Value == 2 {
				firstIndex = i + 1
			}
			if packet.List[0].Value == 6 {
				secondIndex = i + 1
			}
			if firstIndex > -1 && secondIndex > -1 {
				break
			}
		}
	}

	if firstIndex == -1 || secondIndex == -1 {
		panic(errors.New("didn't find a divider packet"))
	}

	fmt.Println(firstIndex * secondIndex)
}
