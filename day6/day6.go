package day4

import (
	getInput "adventcode/getInputs"
	"errors"
	"fmt"
	"log"
)

const URL = "https://adventofcode.com/2022/day/6/input"
const PACKETLEN = 4
const MESSAGELEN = 14

func GetResult1() {
	r, err := problem1()
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("Resultado del primer problema del día seis: ", r)
}

func GetResult2() {
	r, err := problem2()
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("Resultado del segundo problema del día seis: ", r)
}

func problem1() (int, error) {
	input, err := getInput.GetInput(URL)
	if err != nil {
		return -1, err
	}
	chars := ""

	for i, j := PACKETLEN, 0; i < len(input); i++ {
		chars = input[j:i]
		if checkMarker(chars) {
			return i, nil
		}
		j++
	}
	return -1, errors.New("marker don´t finded")
}

func problem2() (int, error) {
	input, err := getInput.GetInput(URL)
	if err != nil {
		return -1, err
	}
	chars := ""

	for i, j := MESSAGELEN, 0; i < len(input); i++ {
		chars = input[j:i]
		if checkMarker(chars) {
			return i, nil
		}
		j++
	}
	return -1, errors.New("marker don´t finded")
}

func checkMarker(s string) bool {
	for i := 0; i < len(s); i++ {
		for j := len(s) - 1; j > i; j-- {
			if s[i] == s[j] {
				return false
			}
		}
	}
	return true
}
