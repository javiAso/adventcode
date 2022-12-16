package day5

import (
	//getInput "adventcode/getInputs"
	getInput "adventcode/getInputs"
	"fmt"
	"log"
)

const URL = "https://adventofcode.com/2022/day/5/input"

/* "[M]                     [N] [Z]"
"[F]             [R] [Z] [C] [C]"
"[C]     [V]     [L] [N] [G] [V]"
"[W]     [L]     [T] [H] [V] [F] [H]"
"[T]     [T] [W] [F] [B] [P] [J] [L]"
"[D] [L] [H] [J] [C] [G] [S] [R] [M]"
"[L] [B] [C] [P] [S] [D] [M] [Q] [P]"
"[B] [N] [J] [S] [Z] [W] [F] [W] [R]" */

func GetResult1() {
	r, err := problem1()
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("Resultado del primer problema del día cinco: ", r)
}

func problem1() (int, error) {
	array := make([][]string, 0)
	//arrayAux := make([]string,0)
	//array de ejemplo
	array = append(array, []string{"Z", "N"})
	array = append(array, []string{"M", "C", "D"})
	array = append(array, []string{"P"})

	fmt.Println(array)

	toMove := 1
	moveFrom := 1
	moveTo := 0
	var cratesRemoved []string
	array[moveFrom], cratesRemoved = removeCratesFromStack(array[moveFrom], toMove)
	fmt.Println(array)
	fmt.Println(cratesRemoved)
	array[moveTo] = addCratesToStack(array[moveTo], cratesRemoved)
	fmt.Println(array)
	//movimiento primero

	input, err := getInput.GetInput(URL)
	if err != nil {
		return 0, err
	}
	//fmt.Println(input)
	fmt.Println(getInitialStacks(input))

	return 0, nil
}

func removeCratesFromStack(stack []string, quantity int) ([]string, []string) {

	return stack[0 : len(stack)-quantity], stack[len(stack)-quantity:]

}

func addCratesToStack(stack []string, crates []string) []string {

	return append(stack, crates...)

}

/* func fillStacks (input string)([][]string){

	for i := 0; i < len(input); i++ {
		if input[i] == 10  && input[i-1] == 10{
			return input[0:i]
		}
	}
} */

func getInitialStacks(input string) string {
	for i := 0; i < len(input); i++ {
		if input[i] == 10 && input[i-1] == 10 {
			return input[0:i]
		}
	}
	return ""
}
