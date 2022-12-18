package day5

import (
	getInput "adventcode/getInputs"
	"errors"
	"fmt"
	"log"
	"strconv"
)

const URL = "https://adventofcode.com/2022/day/5/input"

func GetResult1() {
	r, err := problem1()
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("Resultado del primer problema del día cinco: ", r)
}

func problem1() (string, error) {
	array := make([][]string, 9)

	input, err := getInput.GetInput(URL)
	if err != nil {
		return "", err
	}

	stacks, orders := getInitialStacksAndOrders(input)

	initStacks(array, stacks)

	executeMovements(orders, array)
	return getTops(array), nil
}

func executeMovements(orders string, stacks [][]string) ([][]string, error) {

	for i := 0; i < len(orders); {
		order, index, err := getOrder(orders, i)
		if err != nil {
			return nil, err
		}
		splitResult, err := myCustomSplit(order, " ")
		if err != nil {
			return nil, err
		}
		toMove, err := strconv.ParseInt(splitResult[1], 10, 64)
		if err != nil {
			return nil, err
		}
		moveFrom, err := strconv.ParseInt(splitResult[3], 10, 64)
		if err != nil {
			return nil, err
		}
		moveTo, err := strconv.ParseInt(splitResult[5], 10, 64)
		if err != nil {
			return nil, err
		}
		stacks = moveCrates(stacks, int(toMove), int(moveFrom)-1, int(moveTo)-1)
		i = index
	}

	return stacks, nil

}

func removeCratesFromStack(stack []string, quantity int) ([]string, []string) {

	return stack[0 : len(stack)-quantity], stack[len(stack)-quantity:]

}

func addCratesToStack(stack []string, crates []string) []string { // Problem 1
	var temp []string
	for i := len(crates) - 1; i > -1; i-- {
		temp = append(temp, crates[i])
	}
	return append(stack, temp...)

}

func addCratesToStack9001(stack []string, crates []string) []string { // Problem 2

	return append(stack, crates...)

}

func getOrder(orders string, index int) (string, int, error) {
	for i := index; i < len(orders); i++ {
		if orders[i] == 10 {
			return orders[index:i], i + 1, nil
		}
	}
	return "", -1, errors.New("no order finded)")
}

func getInitialStacksAndOrders(input string) (string, string) {
	for i := 0; i < len(input); i++ {
		if input[i] == 10 && input[i-1] == 10 {
			return input[0 : i-1], input[i+1:]
		}
	}
	return "", ""
}

func moveCrates(array [][]string, toMove int, moveFrom int, moveTo int) [][]string {

	var cratesRemoved []string
	array[moveFrom], cratesRemoved = removeCratesFromStack(array[moveFrom], toMove)
	array[moveTo] = addCratesToStack9001(array[moveTo], cratesRemoved)
	return array
}

func myCustomSplit(s string, c string) ([]string, error) {

	t := ""
	a := make([]string, 0)
	j := 0

	for i := 0; i < len(s); i++ {
		if s[i] == c[0] {
			a = append(a, s[j:i])
			j = i + 1
		} else {
			t += string(s[i])
		}
	}
	return append(a, s[j:]), nil
}

func initStacks(array [][]string, stacks string) [][]string {
	s := ""
	for i := 0; i < len(stacks); i++ {
		if stacks[i] == 10 {
			for j, k := 1, 0; j < len(s); j += 4 {
				if s[j] != 32 {
					array[k] = append(array[k], string(s[j]))
				}
				k++
			}
			s = ""
		} else {
			s += string(stacks[i])
		}
	}

	return flipStacks(array)
}

func flipStacks(array [][]string) [][]string {

	for i := 0; i < len(array); i++ {
		var tempArray []string
		for k := len(array[i]) - 1; k > -1; k-- {
			tempArray = append(tempArray, array[i][k])
		}
		array[i] = tempArray
	}

	return array

}

func getTops(array [][]string) string {
	s := ""
	for i := 0; i < len(array); i++ {
		s += array[i][len(array[i])-1]
	}
	return s
}
