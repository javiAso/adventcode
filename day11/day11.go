package day11

import (
	"fmt"
	"log"
	"strconv"
)

const URL = "https://adventofcode.com/2022/day/11/input"
const SALTO = 10

var monkeys = make([]Monkey, 0)
var sample string = "Monkey 0:\n" +
	"Starting items: 79, 98\n" +
	"Operation: new = old * 19\n" +
	"Test: divisible by 23\n" +
	"If true: throw to monkey 2\n" +
	"If false: throw to monkey 3\n" +
	"\n" +
	"Monkey 1:\n" +
	"Starting items: 54, 65, 75, 74\n" +
	"Operation: new = old + 6\n" +
	"Test: divisible by 19\n" +
	"If true: throw to monkey 2\n" +
	"If false: throw to monkey 0\n" +
	"\n" +
	"Monkey 2:\n" +
	"Starting items: 79, 60, 97\n" +
	"Operation: new = old * old\n" +
	"Test: divisible by 13\n" +
	"If true: throw to monkey 1\n" +
	"If false: throw to monkey 3\n" +
	"\n" +
	"Monkey 3:\n" +
	"Starting items: 74\n" +
	"Operation: new = old + 3\n" +
	"Test: divisible by 17\n" +
	"If true: throw to monkey 0\n" +
	"If false: throw to monkey 1\n"

func GetResult1() {
	r, err := problem1()
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("El resultado del primer problema del dia 10 es: ", r)
}
func problem1() (int, error) {
	/* 	input, err := getInput.GetInput(URL)
	   	if err != nil {
	   		return -1, err
	   	} */
	s := ""

	for i := 0; i < len(sample); i++ {
		if sample[i] == SALTO {
			s += "\n"
			if sample[i-1] == SALTO {
				createMonkey(s)
				s = ""
			}

		} else {
			s += string(sample[i])
		}
	}
	createMonkey(s)
	return 0, nil
}

func createMonkey(monkey string) Monkey {
	m := Monkey{}
	monkey = monkey[:len(monkey)-1]
	s := ""
	for i, j := 0, 0; i < len(monkey); i++ {
		if monkey[i] == SALTO {
			fmt.Print(s)
			processSentence(j, s, &m)
			s = ""
			j++

		} else {
			s += string(monkey[i])
		}
	}
	return m
}

func processSentence(j int, sent string, m *Monkey) error {
	switch j {
	case 0:
		sent = sent[:len(sent)-1]
		id, err := strconv.ParseUint(myCustomSplit(sent, " ")[1], 10, 8)
		if err != nil {
			return err
		}
		m.Id = uint8(id)
		return nil
	case 3:
		div, err := strconv.ParseInt(myCustomSplit(sent, " ")[3], 10, 64)
		if err != nil {
			return err
		}
		m.Test = int(div)
		return nil
	}
	return nil
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
