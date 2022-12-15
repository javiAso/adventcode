package day2

import (
	getInput "adventcode/getInputs"
	"errors"
	"fmt"
	"log"
)

const URL = "https://adventofcode.com/2022/day/2/input"

const ROCK = "A"
const PAPER = "B"
const SCISSORS = "C"

const MYROCK = "X"
const MYPAPER = "Y"
const MYSCISSORS = "Z"

const ROCKPOINTS = 1
const PAPERPOINTS = 2
const SCISSORSPOINTS = 3

const LOSSPOINTS = 0
const DRAWPOINTS = 3
const WINPOINTS = 6

func GetResult1() {
	r, err := problem1()
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("Resultado del primer problema del día dos: ", r)
}

func GetResult2() {
	r, err := problem2()
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("Resultado del segundo problema del día dos: ", r)
}

func problem1() (int, error) {
	var totalPoints int
	input, err := getInput.GetInput(URL)
	if err != nil {
		return 0, err
	}

	for i := 0; i < len(input); i += 4 {
		totalPointsTemp, err := rockPaperScissors(string(input[i]), string(input[i+2]))
		if err != nil {
			return 0, err
		}
		totalPoints += totalPointsTemp
	}
	return totalPoints, nil
}

func problem2() (int, error) {
	var totalPoints int
	input, err := getInput.GetInput(URL)
	if err != nil {
		return 0, err
	}

	for i := 0; i < len(input); i += 4 {
		totalPointsTemp, err := rockPaperScissors2(string(input[i]), string(input[i+2]))
		if err != nil {
			return 0, err
		}
		totalPoints += totalPointsTemp
	}
	return totalPoints, nil
}

func rockPaperScissors(shape1 string, shape2 string) (int, error) {
	if shape1 == ROCK {
		if shape2 == MYSCISSORS { // Loss
			return LOSSPOINTS + SCISSORSPOINTS, nil
		}
		if shape2 == MYROCK { // Draw
			return DRAWPOINTS + ROCKPOINTS, nil
		}
		if shape2 == MYPAPER { // Win
			return WINPOINTS + PAPERPOINTS, nil
		}
	}
	if shape1 == PAPER {
		if shape2 == MYROCK { // Loss
			return LOSSPOINTS + ROCKPOINTS, nil
		}
		if shape2 == MYPAPER { // Draw
			return DRAWPOINTS + PAPERPOINTS, nil
		}
		if shape2 == MYSCISSORS { // Win
			return WINPOINTS + SCISSORSPOINTS, nil
		}
	}
	if shape1 == SCISSORS {
		if shape2 == MYPAPER { // Loss
			return LOSSPOINTS + PAPERPOINTS, nil
		}
		if shape2 == MYSCISSORS { // Draw
			return DRAWPOINTS + SCISSORSPOINTS, nil
		}
		if shape2 == MYROCK { // Win
			return WINPOINTS + ROCKPOINTS, nil
		}
	}
	return 0, errors.New("someone not play a valid shape")
}

func rockPaperScissors2(shape1 string, shape2 string) (int, error) {

	if shape2 == MYSCISSORS { // Win
		if shape1 == ROCK {
			return WINPOINTS + PAPERPOINTS, nil // Win and choose paper to defeat rock
		}
		if shape1 == PAPER {
			return WINPOINTS + SCISSORSPOINTS, nil // Win and choose scissors to defeat paper
		}
		if shape1 == SCISSORS {
			return WINPOINTS + ROCKPOINTS, nil // Win and choose rock to defeat scissors
		}
	}
	if shape2 == MYROCK { // Loss
		if shape1 == ROCK {
			return LOSSPOINTS + SCISSORSPOINTS, nil // Loss and choose scissors to be defeated by rock
		}
		if shape1 == PAPER {
			return LOSSPOINTS + ROCKPOINTS, nil // Loss and choose rock to be defeated by paper
		}
		if shape1 == SCISSORS {
			return LOSSPOINTS + PAPERPOINTS, nil // Loss and choose paper to be defeated by scissors
		}
	}
	if shape2 == MYPAPER { // Draw
		if shape1 == ROCK {
			return DRAWPOINTS + ROCKPOINTS, nil // Choose rock to draw
		}
		if shape1 == PAPER {
			return DRAWPOINTS + PAPERPOINTS, nil // Choose paper to draw
		}
		if shape1 == SCISSORS {
			return DRAWPOINTS + SCISSORSPOINTS, nil // Choose scissors to draw
		}
	}
	return 0, errors.New("someone not play a valid shape")
}
