package day11

import (
	"advent-of-code-22/lib"
	"fmt"
	"regexp"
	"strings"
)

func Run() {
	input, err := lib.ReadLines("day11", "monkey-business-test.txt") // monkey-business.txt
	lib.HandleError(err)

	loadFile(input)
}

func loadFile(input []string) []Monkey {
	monkeys := []Monkey{}
	var m Monkey

	for _, s := range input {
		switch {
		case strings.HasPrefix(s, "Monkey"):
			m = NewMonkey(strings.TrimSuffix(s, ":"))
		case strings.HasPrefix(s, "Starting"):
			data := strings.SplitN(s, ":", 2)
			m.Items = strToIntArray(data[1])
		case strings.HasPrefix(s, "Operation"):
			data := strings.SplitN(s, ":", 2)
			m.SetOperation(data[1])
		case strings.HasPrefix(s, "Test"):
			data := strings.SplitN(s, ":", 2)
			m.Test = data[1]
		case strings.HasPrefix(s, "If"):
			data := strings.SplitN(s, ":", 2)

			if strings.Contains(data[0], "true") {
				m.ResultIfTrue = data[1]
			} else {
				m.ResultIfFalse = data[1]
			}
		case len(s) <= 1:
			fmt.Printf("%+v\n", m)
			monkeys = append(monkeys, m)
		}
	}
	fmt.Printf("%+v\n", m)
	monkeys = append(monkeys, m)

	return monkeys
}

func strToIntArray(s string) []int {
	data := strings.Split(s, ",")
	arr := []int{}

	for _, d := range data {
		i := lib.StrToInt(d)
		arr = append(arr, i)
	}

	return arr
}

type Monkey struct {
	Name          string
	Items         []int
	Operation     OperatorFunc
	Test          string
	ResultIfTrue  string
	ResultIfFalse string
}

type OperatorFunc func(int) int

func NewMonkey(name string) Monkey {
	items := []int{}
	testCondition := ""
	resultIfTrue := ""
	resultIfFalse := ""

	return Monkey{name, items, nil, testCondition, resultIfTrue, resultIfFalse}
}

func (m *Monkey) SetOperation(input string) {
	pat := regexp.MustCompile(`new = old ([+*-/]) (\w+|\d+)`)
	vals := pat.FindAllStringSubmatch(input, -1)
	operator, operand := vals[0][1], vals[0][2]

	var f OperatorFunc
	if operand == "old" {
		switch operator {
		case "+":
			f = func(old int) int { return old + old }
		case "-":
			f = func(old int) int { return 0 }
		case "*":
			f = func(old int) int { return old * old }
		case "/":
			f = func(old int) int { return 1 }
		}
	} else {
		op := lib.StrToInt(operand)
		switch operator {
		case "+":
			f = func(old int) int { return old + op }
		case "-":
			f = func(old int) int { return old - op }
		case "*":
			f = func(old int) int { return old * op }
		case "/":
			f = func(old int) int { return old / op }
		}
	}

	m.Operation = f
}
