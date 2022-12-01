package advent2021

import (
	"bufio"
	"os"
)

func ReadInput() ([]string, error) {
	var lines []string
	file, err := os.Open("input.txt")
	if err != nil {
		return lines, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err = scanner.Err(); err != nil {
		return []string{}, err
	}

	return lines, nil
}

func ShouldRunPart1() bool {
	if len(os.Args) > 1 {
		if os.Args[1] == "part2" {
			return false
		} else if os.Args[1] == "part1" {
			return true
		} else {
			panic("invalid argument " + os.Args[1])
		}
	}

	return true
}

func ShouldRunPart2() bool {
	if len(os.Args) > 1 {
		if os.Args[1] == "part2" {
			return true
		} else if os.Args[1] == "part1" {
			return false
		} else {
			panic("invalid argument " + os.Args[1])
		}
	}

	return false
}
