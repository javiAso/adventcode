package day11

import (
	"fmt"
	"strconv"
)

type Monkey struct {
	Id        uint8
	Items     []int
	Operation string
	Test      int
	True      uint8
	False     uint8
}

func PrintMonkey(m Monkey) {
	fmt.Println("Monkey " + strconv.Itoa(int(m.Id)))
	items := "Items: "
	for i := 0; i < len(m.Items); i++ {
		items += strconv.Itoa(m.Items[i]) + ", "
	}
	fmt.Println(items[:len(items)-1])
	fmt.Println(m.Operation)
	fmt.Println("Test: divisible by ", m.Test)
	fmt.Println("  If true: throw to monkey ", m.True)
	fmt.Println("  If false: throw to monkey ", m.False)
}
