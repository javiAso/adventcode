package day10

import (
	getInput "adventcode/getInputs"
	"fmt"
	"log"
	"strconv"
)

const URL = "https://adventofcode.com/2022/day/10/input"
const SALTO = 10

var sample = "addx 15\n" +
	"addx -11\n" +
	"addx 6\n" +
	"addx -3\n" +
	"addx 5\n" +
	"addx -1\n" +
	"addx -8\n" +
	"addx 13\n" +
	"addx 4\n" +
	"noop\n" +
	"addx -1\n" +
	"addx 5\n" +
	"addx -1\n" +
	"addx 5\n" +
	"addx -1\n" +
	"addx 5\n" +
	"addx -1\n" +
	"addx 5\n" +
	"addx -1\n" +
	"addx -35\n" +
	"addx 1\n" +
	"addx 24\n" +
	"addx -19\n" +
	"addx 1\n" +
	"addx 16\n" +
	"addx -11\n" +
	"noop\n" +
	"noop\n" +
	"addx 21\n" +
	"addx -15\n" +
	"noop\n" +
	"noop\n" +
	"addx -3\n" +
	"addx 9\n" +
	"addx 1\n" +
	"addx -3\n" +
	"addx 8\n" +
	"addx 1\n" +
	"addx 5\n" +
	"noop\n" +
	"noop\n" +
	"noop\n" +
	"noop\n" +
	"noop\n" +
	"addx -36\n" +
	"noop\n" +
	"addx 1\n" +
	"addx 7\n" +
	"noop\n" +
	"noop\n" +
	"noop\n" +
	"addx 2\n" +
	"addx 6\n" +
	"noop\n" +
	"noop\n" +
	"noop\n" +
	"noop\n" +
	"noop\n" +
	"addx 1\n" +
	"noop\n" +
	"noop\n" +
	"addx 7\n" +
	"addx 1\n" +
	"noop\n" +
	"addx -13\n" +
	"addx 13\n" +
	"addx 7\n" +
	"noop\n" +
	"addx 1\n" +
	"addx -33\n" +
	"noop\n" +
	"noop\n" +
	"noop\n" +
	"addx 2\n" +
	"noop\n" +
	"noop\n" +
	"noop\n" +
	"addx 8\n" +
	"noop\n" +
	"addx -1\n" +
	"addx 2\n" +
	"addx 1\n" +
	"noop\n" +
	"addx 17\n" +
	"addx -9\n" +
	"addx 1\n" +
	"addx 1\n" +
	"addx -3\n" +
	"addx 11\n" +
	"noop\n" +
	"noop\n" +
	"addx 1\n" +
	"noop\n" +
	"addx 1\n" +
	"noop\n" +
	"noop\n" +
	"addx -13\n" +
	"addx -19\n" +
	"addx 1\n" +
	"addx 3\n" +
	"addx 26\n" +
	"addx -30\n" +
	"addx 12\n" +
	"addx -1\n" +
	"addx 3\n" +
	"addx 1\n" +
	"noop\n" +
	"noop\n" +
	"noop\n" +
	"addx -9\n" +
	"addx 18\n" +
	"addx 1\n" +
	"addx 2\n" +
	"noop\n" +
	"noop\n" +
	"addx 9\n" +
	"noop\n" +
	"noop\n" +
	"noop\n" +
	"addx -1\n" +
	"addx 2\n" +
	"addx -37\n" +
	"addx 1\n" +
	"addx 3\n" +
	"noop\n" +
	"addx 15\n" +
	"addx -21\n" +
	"addx 22\n" +
	"addx -6\n" +
	"addx 1\n" +
	"noop\n" +
	"addx 2\n" +
	"addx 1\n" +
	"noop\n" +
	"addx -10\n" +
	"noop\n" +
	"noop\n" +
	"addx 20\n" +
	"addx 1\n" +
	"addx 2\n" +
	"addx 2\n" +
	"addx -6\n" +
	"addx -11\n" +
	"noop\n" +
	"noop\n" +
	"noop\n"

func GetResult1() {
	r, err := problem1()
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("El resultado del primer problema del dia 10 es: ", r)
}

func GetResult2() {
	fmt.Println("El resultado del segundo problema del dia 10 es: ")
	err := problem2()
	if err != nil {
		log.Fatal(err.Error())
	}
}

func problem1() (int, error) {
	input, err := getInput.GetInput(URL)
	if err != nil {
		return -1, err
	}
	s := ""

	xCycleBuffer := make([]int, 0)
	for i, j := 0, 0; i < len(input); i++ {
		if input[i] == SALTO {
			buffer, err := processCycle(s, xCycleBuffer)
			if err != nil {
				return -1, err
			}
			xCycleBuffer = buffer
			s = ""
			j++
		} else {
			s += string(input[i])
		}
	}
	return sumBuffer(xCycleBuffer), nil
}

func problem2() error {
	input, err := getInput.GetInput(URL)
	if err != nil {
		return err
	}
	s := ""

	xCycleBuffer := make([]int, 0)
	for i, j := 0, 0; i < len(input); i++ {
		if input[i] == SALTO {
			buffer, err := processCycle(s, xCycleBuffer)
			if err != nil {
				return err
			}
			xCycleBuffer = buffer
			s = ""
			j++
		} else {
			s += string(input[i])
		}
	}
	printSprite(xCycleBuffer)
	return nil
}

func processCycle(cycle string, buffer []int) ([]int, error) {
	buffer = append(buffer, 0)
	if cycle == "noop" {
		return buffer, nil
	}
	cycleSplit := myCustomSplit(cycle, " ")
	addValue, err := strconv.ParseInt(cycleSplit[1], 10, 64)
	if err != nil {
		return nil, err
	}
	return append(buffer, int(addValue)), nil
}

func sumBuffer(buffer []int) int {
	xRegisterValue := 1
	var xSum int

	for j := 0; j < len(buffer); j++ {
		if j+1 == 20 || j+1 == 60 || j+1 == 100 || j+1 == 140 || j+1 == 180 || j+1 == 220 {
			xSum += xRegisterValue * (j + 1)
		}

		xRegisterValue += buffer[j]
	}
	return xSum
}

func printSprite(buffer []int) {
	sprite := ""
	xRegisterValue := 1
	for j, cycle := 0, 0; j < len(buffer); j++ {

		if cycle < xRegisterValue-1 || cycle > xRegisterValue+1 {
			sprite += "."
		} else {
			sprite += "#"
		}

		cycle++
		if cycle == 40 {
			cycle = 0
			sprite += "\n"
		}
		xRegisterValue += buffer[j]
	}
	fmt.Print(sprite)
}

func myCustomSplit(s string, c string) []string {
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
	return append(a, s[j:])
}
