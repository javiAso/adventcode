package day11

import (
	"fmt"
	"strconv"
)

const MULT = 42
const SUM = 43

type Monkey struct {
	Id             uint8
	Items          []int
	Operation      string
	Test           int
	True           uint8
	False          uint8
	ItemsInspected int
}

func printMonkeys(m []Monkey) {

	for i := 0; i < len(m); i++ {
		printMonkey(m[i])
		fmt.Println()
	}

}

func printMonkey(m Monkey) {
	fmt.Println("Monkey " + strconv.Itoa(int(m.Id)) + ":")
	items := "Items: "
	for i := 0; i < len(m.Items); i++ {
		items += strconv.Itoa(m.Items[i]) + ", "
	}
	fmt.Println(items[:len(items)-2])
	fmt.Println("Operation: new = " + m.Operation)
	fmt.Println("Test: divisible by", m.Test)
	fmt.Println("  If true: throw to monkey", m.True)
	fmt.Println("  If false: throw to monkey", m.False)
	fmt.Println("  Items inspected", m.ItemsInspected)
}

func inspectItem(m *Monkey, i int) error {
	split := myCustomSplit(m.Operation, " ")
	op := split[1][0]
	if split[2] == "old" {
		m.Items[i] = doOperation(op, m.Items[i], m.Items[i])
	} else {
		operand2, err := strconv.ParseInt(split[2], 10, 64)
		if err != nil {
			return err
		}
		m.Items[i] = doOperation(op, m.Items[i], int(operand2))
	}
	m.ItemsInspected++
	return nil
}

func doOperation(op byte, op1 int, op2 int) int {
	switch op {
	case SUM:
		return op1 + op2
	case MULT:
		return op1 * op2
	default:
		return 0
	}
}

func getBored(m Monkey, i int) {
	m.Items[i] /= 3
}

func throwItem(m *Monkey, i int, ms []Monkey) {
	if m.Items[i]%m.Test == 0 {
		ms[m.True].Items = append(ms[m.True].Items, m.Items[i])
	} else {
		ms[m.False].Items = append(ms[m.False].Items, m.Items[i])
	}
	m.Items = append(m.Items[:i], m.Items[i+1:]...)
	ms[m.Id] = *m
}

func monkeysGetFun(ms []Monkey) {
	for j := 0; j < 20; j++ {
		for i := 0; i < len(ms); i++ {
			monkeyGetFun(ms[i], ms)
		}
	}
}

func monkeyGetFun(m Monkey, ms []Monkey) {
	for i := 0; i < len(m.Items); {
		inspectItem(&m, i)
		getBored(m, i)
		throwItem(&m, i, ms)
	}
}

func getMonkeyBusiness(ms []Monkey) int {
	m1 := ms[0]
	m2 := ms[1]
	for i := 2; i < len(ms); i++ {
		if ms[i].ItemsInspected > m1.ItemsInspected {
			aux := m1
			m1 = ms[i]
			if aux.ItemsInspected > m2.ItemsInspected {
				m2 = aux
			}
		} else {
			if ms[i].ItemsInspected > m2.ItemsInspected {
				m2 = ms[i]
			}
		}
	}
	return m1.ItemsInspected * m2.ItemsInspected
}
