package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	advent "github.com/BenJuan26/advent2022"
)

func GCD(a, b int) int {
	for b != 0 {
		temp := b
		b = a % b
		a = temp
	}
	return a
}

func LCM(a, b int, extra ...int) int {
	result := a * b / GCD(a, b)

	for _, c := range extra {
		result = LCM(result, c)
	}

	return result
}

func Part2() {
	lines, err := advent.ReadInput()
	if err != nil {
		panic(err)
	}

	stages := []string{"MONKEY", "ITEMS", "OPERATION", "TEST", "TRUE", "FALSE", "BLANK"}
	stage := 0

	monkeys := []*Monkey{}
	currentMonkey := &Monkey{}
	divisors := []int{}
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
			divisors = append(divisors, int(currentMonkey.TestDivisor))
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

	lcm := LCM(divisors[0], divisors[1], divisors[2:]...)

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
				monkey.Items[i] %= lcm
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
