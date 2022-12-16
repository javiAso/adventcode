package day4

import (
	getInput "adventcode/getInputs"
	"errors"
	"fmt"
	"log"
	"strconv"
)

const URL = "https://adventofcode.com/2022/day/4/input"
const COMA = 44
const SALTO = 10
const GUION = 45

func GetResult1() {
	r, err := problem1()
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("Resultado del primer problema del día cuatro: ", r)
}

func GetResult2() {
	r, err := problem2()
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("Resultado del segundo problema del día cuatro: ", r)
}

func problem1() (int, error) {
	var assignmentPair string = ""
	var result int
	input, err := getInput.GetInput(URL)
	if err != nil {
		return 0, err
	}
	for i := 0; i < len(input); i++ {
		if input[i] == SALTO {
			assignment1, assignment2, err := separeAssignments(assignmentPair)
			if err != nil {
				return 0, err
			}
			index1assigment1, index2assigment1, err := getIndex(assignment1)
			if err != nil {
				return 0, err
			}
			index1assigment2, index2assigment2, err := getIndex(assignment2)
			if err != nil {
				return 0, err
			}

			fullyContained := compareIndex(index1assigment1, index2assigment1, index1assigment2, index2assigment2)
			assignmentPair = ""
			if fullyContained {
				result++
			}

		} else {
			assignmentPair += string(input[i])
		}
	}
	return result, nil
}

func problem2() (int, error) {
	var assignmentPair string = ""
	var result int
	input, err := getInput.GetInput(URL)
	if err != nil {
		return 0, err
	}
	for i := 0; i < len(input); i++ {
		if input[i] == SALTO {
			assignment1, assignment2, err := separeAssignments(assignmentPair)
			if err != nil {
				return 0, err
			}
			index1assigment1, index2assigment1, err := getIndex(assignment1)
			if err != nil {
				return 0, err
			}
			index1assigment2, index2assigment2, err := getIndex(assignment2)
			if err != nil {
				return 0, err
			}

			fullyContained := getOverlap(index1assigment1, index2assigment1, index1assigment2, index2assigment2)
			assignmentPair = ""
			if fullyContained {
				result++
			}

		} else {
			assignmentPair += string(input[i])
		}
	}
	return result, nil
}

func separeAssignments(assignmentPair string) (string, string, error) {
	var assignment1 string = ""
	var assignment2 string = ""
	comaIndex, err := getStringIndex(assignmentPair, ",")
	if err != nil {
		return "", "", err
	}
	for i := 0; i < comaIndex; i++ {
		assignment1 += string(assignmentPair[i])
	}

	for i := comaIndex + 1; i < len(assignmentPair); i++ {
		assignment2 += string(assignmentPair[i])
	}

	return assignment1, assignment2, nil

}

func getIndex(a string) (int, int, error) {
	var index1 = ""
	var index2 = ""
	dashPosition, err := getStringIndex(a, "-")
	if err != nil {
		return -1, -1, err
	}
	for i := 0; i < dashPosition; i++ {
		index1 += string(a[i])
	}
	for i := dashPosition + 1; i < len(a); i++ {
		index2 += string(a[i])
	}
	index1Int, err := strconv.ParseInt(index1, 10, 64)
	if err != nil {
		return -1, -1, err
	}
	index2Int, err := strconv.ParseInt(index2, 10, 64)
	if err != nil {
		return -1, -1, err
	}
	return int(index1Int), int(index2Int), nil
}

func getStringIndex(s string, c string) (int, error) {
	if len(c) != 1 {
		return -1, errors.New("you give me more than one char ;)")
	}
	for i := 0; i < len(s); i++ {
		if s[i] == c[0] {
			return i, nil
		}
	}
	return -1, nil
}

func compareIndex(i1a1 int, i2a1 int, i1a2 int, i2a2 int) bool {
	return (i1a1 <= i1a2 && i2a1 >= i2a2) || (i1a2 <= i1a1 && i2a2 >= i2a1)
}

func getOverlap(i1a1 int, i2a1 int, i1a2 int, i2a2 int) bool {
	return !(i2a2 < i1a1 || i2a1 < i1a2)
}
