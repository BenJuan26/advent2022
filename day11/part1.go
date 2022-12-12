package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	advent "github.com/BenJuan26/advent2022"
)

type Monkey struct {
	Items            []int
	OperationType    string
	OperationValue   string
	TestDivisor      int
	TrueMonkey       int
	FalseMonkey      int
	TotalInspections int
}

func Part1() {
	lines, err := advent.ReadInput()
	if err != nil {
		panic(err)
	}

	stages := []string{"MONKEY", "ITEMS", "OPERATION", "TEST", "TRUE", "FALSE", "BLANK"}
	stage := 0

	monkeys := []*Monkey{}
	currentMonkey := &Monkey{}
	for _, line := range lines {
		switch stage {
		case 1:
			fields := strings.Split(line, ": ")
			itemFields := strings.Split(fields[1], ", ")
			for _, itemString := range itemFields {
				item, err := strconv.Atoi(itemString)
				if err != nil {
					panic(err)
				}
				currentMonkey.Items = append(currentMonkey.Items, item)
			}
		case 2:
			fields := strings.Split(line, "old ")
			operationFields := strings.Split(fields[1], " ")
			currentMonkey.OperationType = operationFields[0]
			currentMonkey.OperationValue = operationFields[1]
			if err != nil {
				panic(err)
			}
		case 3:
			fields := strings.Split(line, "by ")
			temp, err := strconv.Atoi(fields[1])
			if err != nil {
				panic(err)
			}
			currentMonkey.TestDivisor = temp
		case 4:
			fields := strings.Split(line, "monkey ")
			currentMonkey.TrueMonkey, err = strconv.Atoi(fields[1])
			if err != nil {
				panic(err)
			}
		case 5:
			fields := strings.Split(line, "monkey ")
			currentMonkey.FalseMonkey, err = strconv.Atoi(fields[1])
			if err != nil {
				panic(err)
			}
		case 6:
			monkeys = append(monkeys, currentMonkey)
			currentMonkey = &Monkey{}
		}
		stage += 1
		stage = stage % len(stages)
	}
	monkeys = append(monkeys, currentMonkey)

	for round := 0; round < 10000; round++ {
		for _, monkey := range monkeys {
			for i := range monkey.Items {
				monkey.TotalInspections += 1
				if monkey.OperationType == "+" {
					num, err := strconv.Atoi(monkey.OperationValue)
					if err != nil {
						panic(err)
					}
					monkey.Items[i] += num
				} else { // multiplication
					if monkey.OperationValue == "old" {
						monkey.Items[i] *= monkey.Items[i]
					} else {
						num, err := strconv.Atoi(monkey.OperationValue)
						if err != nil {
							panic(err)
						}
						monkey.Items[i] *= num
					}
				}
				monkey.Items[i] /= 3
				if monkey.Items[i]%monkey.TestDivisor == 0 {
					monkeys[monkey.TrueMonkey].Items = append(monkeys[monkey.TrueMonkey].Items, monkey.Items[i])
				} else {
					monkeys[monkey.FalseMonkey].Items = append(monkeys[monkey.FalseMonkey].Items, monkey.Items[i])
				}
			}
			monkey.Items = []int{}
		}
	}

	inspects := []int{}
	for _, monkey := range monkeys {
		inspects = append(inspects, monkey.TotalInspections)
	}

	sort.Ints(inspects)

	fmt.Println(inspects[len(inspects)-1] * inspects[len(inspects)-2])
}

func main() {
	if advent.ShouldRunPart1() {
		Part1()
	}

	if advent.ShouldRunPart2() {
		Part2()
	}
}
