package day7

import (
	"adventcode/day7/structs"
	getInput "adventcode/getInputs"
	"fmt"
	"log"

	"strconv"
)

const URL = "https://adventofcode.com/2022/day/7/input"
const SALTO = 10
const TOTALDISK = 70000000
const TOTALUPDATE = 30000000

var rootDirectory structs.Directory

func GetResult1() {
	r, err := problem1()
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("El resultado del primer problema del dia 7 es: ", r)
}

func GetResult2() {
	r, err := problem2()
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("El resultado del segundo problema del dia 7 es: ", r)
}

func problem1() (int, error) {

	rootDirectory.Name = "/"
	var dirsToErase structs.Directory
	dirsToErase.Name = "Dirs To Erase"
	input, err := getInput.GetInput(URL)
	if err != nil {
		return -1, err
	}

	err = executeOrders(0, &rootDirectory, input)
	if err != nil {
		return -1, err
	}

	structs.CalculateTotalSizes(&rootDirectory)
	structs.PrintDirectory(rootDirectory, 0)
	structs.GetDirsToErase(rootDirectory, &dirsToErase)
	return getTotal(dirsToErase), nil
}

func problem2() (int, error) {

	rootDirectory.Name = "/"
	var dirsSpaceed structs.Directory
	input, err := getInput.GetInput(URL)
	if err != nil {
		return -1, err
	}

	err = executeOrders(0, &rootDirectory, input)
	if err != nil {
		return -1, err
	}

	structs.CalculateTotalSizes(&rootDirectory)
	structs.GetDirsFreeSpaceEnough(rootDirectory, &dirsSpaceed, TOTALDISK-TOTALUPDATE-rootDirectory.TotalSize)

	return getDirToEraseSize(dirsSpaceed), nil
}

func getDirToEraseSize(d structs.Directory) int {
	size := d.Directorys[0].TotalSize

	for i := 0; i < len(d.Directorys); i++ {
		if d.Directorys[i].TotalSize < size {
			size = d.Directorys[i].TotalSize
		}
	}

	return size
}

func getTotal(d structs.Directory) int {
	var sum int
	for i := 0; i < len(d.Directorys); i++ {
		sum += d.Directorys[i].TotalSize
	}
	return sum
}

func executeOrders(startIndex int, d *structs.Directory, input string) error {
	order := ""
	for i := startIndex; i < len(input); i++ {

		if input[i] == SALTO {
			dirPointer, err := executeOrder(d, order)
			if err != nil {
				return err
			}
			d = dirPointer
			order = ""

		} else {
			order += string(input[i])
		}
	}
	return nil
}

func executeOrder(d *structs.Directory, o string) (*structs.Directory, error) {
	orders, err := myCustomSplit(o, " ")
	if err != nil {
		return nil, err
	}
	switch orders[0] {
	case "$":
		if orders[1] == "cd" {
			newDir := changeDirectory(d, orders[2])
			return newDir, nil
		}

	case "dir":
		appendDirectory(d, orders)
		if err != nil {
			return nil, err
		}
		return d, nil
	default:
		err := appendFile(d, orders)
		if err != nil {
			return nil, err
		}
		return d, nil
	}
	return d, nil
}

func appendDirectory(d *structs.Directory, orders []string) {

	dir := structs.Directory{
		Name:    orders[1],
		Pointer: d,
	}
	d.Directorys = append(d.Directorys, dir)
}

func appendFile(d *structs.Directory, orders []string) error {
	size, err := strconv.ParseInt(orders[0], 10, 64)
	if err != nil {
		return err
	}
	f := structs.File{
		Name: orders[1],
		Size: int(size),
	}
	d.Files = append(d.Files, f)
	return nil
}

func changeDirectory(d *structs.Directory, toChange string) *structs.Directory {
	if toChange == ".." {
		return d.Pointer
	}
	if toChange == "/" {
		return &rootDirectory
	}
	for i := 0; i < len(d.Directorys); i++ {

		if d.Directorys[i].Name == toChange {
			return &d.Directorys[i]
		}
	}
	return nil
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
