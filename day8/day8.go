package day8

import (
	getInput "adventcode/getInputs"
	"fmt"
	"log"
	"strconv"
)

const URL = "https://adventofcode.com/2022/day/8/input"
const SALTO = 10

func GetResult1() {
	r, err := problem1()
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("El resultado del primer problema del dia 8 es: ", r)
}

func GetResult2() {
	r, err := problem2()
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("El resultado del segundo problema del dia 8 es: ", r)
}

func problem1() (int, error) {
	input, err := getInput.GetInput(URL)
	if err != nil {
		return -1, err
	}
	grid, err := fillGrid(input)
	if err != nil {
		return -1, err
	}
	return inspectGrid(grid), nil
}

func problem2() (int, error) {
	input, err := getInput.GetInput(URL)
	if err != nil {
		return -1, err
	}
	grid, err := fillGrid(input)
	if err != nil {
		return -1, err
	}
	return getHighScenicScore(grid), nil
}

func fillGrid(input string) ([][]uint8, error) {
	grid := make([][]uint8, 1)
	for i, j, k := 0, 0, 0; i < len(input); i++ {
		if input[i] == SALTO {
			k++
			j = 0
			line := make([]uint8, 0)
			grid = append(grid, line)
		} else {
			tree, err := (strconv.ParseInt(string(input[i]), 10, 8))
			if err != nil {
				return grid, err
			}
			grid[k] = append(grid[k], uint8((tree)))
			j++
		}
	}
	return grid[:len(grid)-1], nil
}

func inspectGrid(grid [][]uint8) int {
	visibleTrees := getEdgeTrees(grid)
	for i := 1; i < len(grid)-1; i++ {
		for j := 1; j < len(grid[i])-1; j++ {
			if getTreeVisibility(i, j, grid[i][j], grid) {
				visibleTrees++
			}
		}
	}
	return visibleTrees
}

func getHighScenicScore(grid [][]uint8) int {
	var highScenicScore int
	for i := 1; i < len(grid)-1; i++ {
		for j := 1; j < len(grid[i])-1; j++ {
			treeScenicScore := getTreeScenicScore(i, j, grid[i][j], grid)
			if treeScenicScore > highScenicScore {
				highScenicScore = treeScenicScore
			}
		}
	}
	return highScenicScore
}

func getEdgeTrees(grid [][]uint8) int {
	return len(grid[0])*2 + len(grid)*2 - 4
}

func getTreeVisibility(i int, j int, tree uint8, grid [][]uint8) bool {
	return lookLeft(i, j, tree, grid) || lookRight(i, j, tree, grid) || lookUp(i, j, tree, grid) || lookDown(i, j, tree, grid)
}

func getTreeScenicScore(i int, j int, tree uint8, grid [][]uint8) int {
	return countLeft(i, j, tree, grid) * countRight(i, j, tree, grid) * countUp(i, j, tree, grid) * countDown(i, j, tree, grid)
}

func lookUp(i int, j int, tree uint8, grid [][]uint8) bool {
	for k := i - 1; k > -1; k-- {
		if tree <= grid[k][j] {
			return false
		}
	}
	return true
}

func lookDown(i int, j int, tree uint8, grid [][]uint8) bool {
	for k := i + 1; k < len(grid); k++ {
		if tree <= grid[k][j] {
			return false
		}
	}
	return true
}

func lookRight(i int, j int, tree uint8, grid [][]uint8) bool {
	for k := j + 1; k < len(grid[i]); k++ {
		if tree <= grid[i][k] {
			return false
		}
	}
	return true
}

func lookLeft(i int, j int, tree uint8, grid [][]uint8) bool {
	for k := j - 1; k > -1; k-- {
		if tree <= grid[i][k] {
			return false
		}
	}
	return true
}

func countUp(i int, j int, tree uint8, grid [][]uint8) int {
	var view int
	for k := i - 1; k > -1; k-- {
		view++
		if tree <= grid[k][j] {
			return view
		}
	}
	return view
}

func countDown(i int, j int, tree uint8, grid [][]uint8) int {
	var view int
	for k := i + 1; k < len(grid); k++ {
		view++
		if tree <= grid[k][j] {
			return view
		}
	}
	return view
}

func countRight(i int, j int, tree uint8, grid [][]uint8) int {
	var view int
	for k := j + 1; k < len(grid[i]); k++ {
		view++
		if tree <= grid[i][k] {
			return view
		}
	}
	return view
}

func countLeft(i int, j int, tree uint8, grid [][]uint8) int {
	var view int
	for k := j - 1; k > -1; k-- {
		view++
		if tree <= grid[i][k] {
			return view
		}
	}
	return view
}
