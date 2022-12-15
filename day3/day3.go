package day3

import (
	getInput "adventcode/getInputs"
	"errors"
	"fmt"
	"log"
)

const URL = "https://adventofcode.com/2022/day/3/input"

var priorityArray = []string{"", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

func GetResult1() {
	r, err := problem1()
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("Resultado del primer problema del día tres: ", r)
}

func GetResult2() {
	r, err := problem2()
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("Resultado del segundo problema del día tres: ", r)
}

func problem1() (int, error) {
	var prioritySum int
	var rucksack string = ""
	input, err := getInput.GetInput(URL)
	if err != nil {
		return 0, err
	}

	for i := 0; i < len(input); i++ {
		if input[i] == 10 {
			priority, err := searchRucksack(rucksack)
			if err != nil {
				return 0, err
			}
			prioritySum += priority
			rucksack = ""

		} else {
			rucksack += string(input[i])
		}

	}

	return prioritySum, nil
}

func searchRucksack(rucksack string) (int, error) {
	var compartiment1 string = ""
	var compartiment2 string = ""
	for i := 0; i < len(rucksack)/2; i++ {
		compartiment1 += string(rucksack[i])
	}
	for i := len(rucksack) / 2; i < len(rucksack); i++ {
		compartiment2 += string(rucksack[i])
	}

	return compareCompartiments(compartiment1, compartiment2)
}

func compareCompartiments(c1 string, c2 string) (int, error) {
	for i := 0; i < len(c1); i++ {
		for j := 0; j < len(c2); j++ {
			if c1[i] == c2[j] {
				return searchPriority(string(c1[i]))
			}
		}
	}
	return 0, errors.New("there is a invalid char: ")
}

func searchPriority(x string) (int, error) {
	for i := 0; i < len(priorityArray); i++ {
		if priorityArray[i] == x {
			return i, nil
		}
	}
	return 0, errors.New("there is a invalid char: " + x)
}

func problem2() (int, error) {
	var prioritySum int
	var rucksack string = ""
	var rucksackCounter int
	var rucksackArray = [3]string{"", "", ""}
	input, err := getInput.GetInput(URL)
	if err != nil {
		return 0, err
	}

	for i := 0; i < len(input); i++ {
		if input[i] == 10 {
			rucksackArray[rucksackCounter] = rucksack
			rucksackCounter++
			rucksack = ""
			if rucksackCounter == 3 {
				prioritySumTemp, err := searchBadges(rucksackArray)
				if err != nil {
					return 0, err
				}
				prioritySum += prioritySumTemp
				rucksackCounter = 0
			}

		} else {
			rucksack += string(input[i])
		}

	}

	return prioritySum, nil
}

func searchBadges(rucksackArray [3]string) (int, error) {
	for i := 0; i < len(rucksackArray[0]); i++ {
		for j := 0; j < len(rucksackArray[1]); j++ {
			if rucksackArray[0][i] == rucksackArray[1][j] {
				for k := 0; k < len(rucksackArray[2]); k++ {
					if rucksackArray[1][j] == rucksackArray[2][k] {
						return searchPriority(string(rucksackArray[2][k]))
					}
				}
			}

		}
	}
	return 0, errors.New("there is a group without budget: \n" + rucksackArray[0] + "\n" + rucksackArray[1] + "\n" + rucksackArray[2] + "\n")
}
