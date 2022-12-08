package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	advent "github.com/BenJuan26/advent2022"
)

type node struct {
	Name     string
	Parent   *node
	Children []*node
	Size     int
	IsDir    bool
}

func (n *node) GetSize() int {
	if !n.IsDir {
		return n.Size
	}

	total := 0
	for _, child := range n.Children {
		total += child.GetSize()
	}

	return total
}

func (n *node) TotalBelowThreshold(threshold int) (size int, totalBelow int) {
	for _, child := range n.Children {
		if child.IsDir {
			s, below := child.TotalBelowThreshold(threshold)
			size += s
			totalBelow += below
		} else {
			size += child.Size
		}
	}

	if size <= threshold {
		totalBelow += size
	}

	return
}

func (n *node) SizeOfIdealDirectory(targetSize, currentBest int) (size int, best int) {
	best = currentBest
	if !n.IsDir {
		size = n.Size
		return
	}

	for _, child := range n.Children {
		if !child.IsDir {
			size += child.Size
			continue
		}

		childSize, nextBest := child.SizeOfIdealDirectory(targetSize, best)
		if nextBest < best {
			best = nextBest
		}

		size += childSize
	}

	if size > targetSize && size < best {
		best = size
	}

	return size, best
}

func Part1() {
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

	_, totalBelow := rootNode.TotalBelowThreshold(100000)
	fmt.Println(totalBelow)
}

func main() {
	if advent.ShouldRunPart1() {
		Part1()
	}

	if advent.ShouldRunPart2() {
		Part2()
	}
}
