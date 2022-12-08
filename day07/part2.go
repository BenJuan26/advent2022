package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"

	advent "github.com/BenJuan26/advent2022"
)

func Part2() {
	lines, err := advent.ReadInput()
	if err != nil {
		panic(err)
	}

	cdRegex := regexp.MustCompile(`\$ cd (.+)`)
	rootNode := &node{Name: `/`, IsDir: true}
	currentNode := rootNode
	for _, line := range lines {
		if match := cdRegex.FindStringSubmatch(line); match != nil {
			cd := match[1]
			if cd == `/` {
				currentNode = rootNode
				continue
			}

			if cd == ".." {
				currentNode = currentNode.Parent
				continue
			}

			newNode := &node{Parent: currentNode, Name: cd, IsDir: true}
			currentNode.Children = append(currentNode.Children, newNode)
			currentNode = newNode
		} else if line != `$ ls` {
			// file list
			fields := strings.Split(line, " ")
			sizeOrDir := fields[0]
			name := fields[1]
			if sizeOrDir == "dir" {
				// do nothing for now
				continue
			}

			size, err := strconv.Atoi(sizeOrDir)
			if err != nil {
				panic(err)
			}

			currentNode.Children = append(currentNode.Children, &node{Parent: currentNode, Name: name, Size: size})
		}
	}

	rootSize := rootNode.GetSize()
	totalCapacity := 70000000
	requiredSpace := 30000000
	currentSpace := totalCapacity - rootSize
	toRemove := requiredSpace - currentSpace

	fmt.Printf("To remove: %d\n", toRemove)

	fmt.Println(rootNode.SizeOfIdealDirectory(toRemove, math.MaxInt))
}
