package day9

import (
	getInput "adventcode/getInputs"
	"fmt"
	"log"
)

const URL = "https://adventofcode.com/2022/day/9/input"
const SALTO = 10

func GetResult1() {
	r, err := problem1()
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("El resultado del primer problema del dia 9 es: ", r)
}

func problem1() (int, error) {
	input, err := getInput.GetInput(URL)
	if err != nil {
		return -1, err
	}
	fmt.Print(input)
	return 0, nil
}
