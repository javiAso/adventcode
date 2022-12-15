package day1

import (
	getInput "adventcode/getInputs"
	"fmt"
	"log"
	"strconv"
)

const URL = "https://adventofcode.com/2022/day/1/input"

func GetResult1() {
	r, err := day1()
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("Resultado del primer problema del día uno: ", r)
}

func GetResult2() {
	var r int64
	topThreeCalories, err := day2()
	if err != nil {
		log.Fatal(err.Error())
	}

	for i := 0; i < len(topThreeCalories); i++ {
		r += topThreeCalories[i]
	}

	fmt.Println("Resultado del segundo problema del día uno: ", r)

}

func day1() (lastElfCalories int64, err error) {

	var calorias string = ""
	var sumCalorias int64
	var sumCaloriasTemp int64

	input, err := getInput.GetInput(URL)
	if err != nil {
		return 0, err
	}

	for i := 0; i < len(input); i++ {
		if input[i] == 10 {
			if calorias != "" {
				sumCaloriasTemp, err = strconv.ParseInt(calorias, 10, 64)
				if err != nil {
					return 0, err
				}
				sumCalorias += sumCaloriasTemp
			}
			if input[i-1] == 10 {
				if lastElfCalories < sumCalorias {
					lastElfCalories = sumCalorias
				}
				sumCalorias = 0
			}
			calorias = ""
		} else {
			calorias += string(input[i])
		}
	}

	return lastElfCalories, nil
}

func day2() (topThreeCalories [3]int64, err error) {

	var calorias string = ""
	var sumCalorias int64
	var sumCaloriasTemp int64

	input, err := getInput.GetInput(URL)
	if err != nil {
		return topThreeCalories, err
	}

	for i := 0; i < len(input); i++ {
		if input[i] == 10 {
			if calorias != "" {
				sumCaloriasTemp, err = strconv.ParseInt(calorias, 10, 64)
				if err != nil {
					return topThreeCalories, err
				}
				sumCalorias += sumCaloriasTemp
			}
			if input[i-1] == 10 {
				compareTopCalories(sumCalorias, &topThreeCalories)
				sumCalorias = 0
			}
			calorias = ""
		} else {
			calorias += string(input[i])
		}
	}

	return topThreeCalories, nil
}

func compareTopCalories(lastElfCalories int64, topThreeCalories *[3]int64) {
	if lastElfCalories > topThreeCalories[0] {
		topThreeCalories[0] = lastElfCalories
		ordenateTopCalories(topThreeCalories)
		return
	}
}

//Este metodo burbuja está copiado de: https://awebytes.wordpress.com/2020/04/12/golang-algoritmos-ordenacion-implementados-bases/

func ordenateTopCalories(topThreeCalories *[3]int64) {
	for i := 1; i < len(topThreeCalories); i++ {
		j := i
		for j > 0 && topThreeCalories[j-1] > topThreeCalories[j] {
			swap(j-1, j, &topThreeCalories)
			j--
		}
	}

}
func swap(previo, actual int, puntero_arreglo **[3]int64) {
	arreglo := *puntero_arreglo
	copia := arreglo[actual]
	arreglo[actual] = arreglo[previo]
	arreglo[previo] = copia
}
