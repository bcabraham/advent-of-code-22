package day11

import (
	"advent-of-code-22/lib"
	"fmt"
	"regexp"
	"strings"
)

var (
	debug = false
)

func Run() {
	input, err := lib.ReadLines("day11", "monkey-business-test.txt") // monkey-business.txt
	lib.HandleError(err)

	monkeys := loadFile(input)

	for round := 1; round <= 10000; round++ {
		for _, monkey := range monkeys {
			monkey.DoTurn(monkeys)
		}

		if round == 1 || round == 20 || round%1000 == 0 {
			fmt.Printf("== After Round %d ==\n", round)
			getTopTwoMonkeys(monkeys)
			fmt.Println()
		}
	}

	fmt.Printf("\n\n")
	first, second := getTopTwoMonkeys(monkeys)
	fmt.Printf("Product: %d * %d = %d\n", first, second, first*second)
}

type OperatorFunc func(int) int

type Monkey struct {
	name          string
	items         *lib.Queue[int]
	operation     string
	opFunc        OperatorFunc
	testCondition string
	testValue     int
	resultIfTrue  string
	trueMonkey    int
	resultIfFalse string
	falseMonkey   int
	inspected     int
}

func NewMonkey(name string, startingItems []int, operation, testCondition, resultIfTrue, resultIfFalse string) *Monkey {
	items := lib.NewQueue[int](startingItems)
	opFunc := getOpFunc(operation)

	data := strings.Replace(testCondition, "divisible by ", "", 1)
	testVal := lib.StrToInt(strings.TrimSpace(data))

	data = strings.Replace(resultIfTrue, "throw to monkey ", "", 1)
	trueResult := lib.StrToInt(strings.TrimSpace(data))

	data = strings.Replace(resultIfFalse, "throw to monkey ", "", 1)
	falseResult := lib.StrToInt(strings.TrimSpace(data))

	m := Monkey{name, items, operation, opFunc, testCondition, testVal, resultIfTrue, trueResult, resultIfFalse, falseResult, 0}

	if debug {
		fmt.Printf("new monkey: %s\n", m.String())
	}

	return &m
}

func getOpFunc(operation string) OperatorFunc {
	pat := regexp.MustCompile(`new = old ([+*-/]) (\w+|\d+)`)
	vals := pat.FindAllStringSubmatch(operation, -1)
	operator, operand := vals[0][1], vals[0][2]

	var f OperatorFunc
	if operand == "old" {
		switch operator {
		case "+":
			f = func(old int) int {
				new := old + old
				if debug {
					fmt.Printf("\tWorry level is increased by %d to %d.\n", old, new)
				}
				return new
			}
		case "-":
			f = func(old int) int { return 0 }
		case "*":
			f = func(old int) int {
				new := old * old
				if debug {
					fmt.Printf("\tWorry level is multiplied by %d to %d.\n", old, new)
				}
				return new
			}
		case "/":
			f = func(old int) int { return 1 }
		}
	} else {
		op := lib.StrToInt(operand)
		switch operator {
		case "+":
			f = func(old int) int {
				new := old + op
				if debug {
					fmt.Printf("\tWorry level is increased by %d to %d.\n", op, new)
				}
				return new
			}
		case "-":
			f = func(old int) int {
				new := old - op
				if debug {
					fmt.Printf("\tWorry level is decreased by %d to %d.\n", op, new)
				}
				return new
			}
		case "*":
			f = func(old int) int {
				new := old * op
				if debug {
					fmt.Printf("\tWorry level is multiplied by %d to %d.\n", op, new)
				}
				return new
			}
		case "/":
			f = func(old int) int {
				new := old / op
				if debug {
					fmt.Printf("\tWorry level is divided by %d to %d.\n", op, new)
				}
				return new
			}
		}
	}

	return f
}

func (m *Monkey) String() string {
	return fmt.Sprintf("name: %s, items: %s, operation: %s, testCondition: %s, testValue: %d, resultIfTrue: %s, trueResult: %d, resultIfFalse: %s, falseResult: %d, inspected: %d", m.name, m.items.String(), m.operation, m.testCondition, m.testValue, m.resultIfTrue, m.trueMonkey, m.resultIfFalse, m.falseMonkey, m.inspected)
}

func (m *Monkey) inspect(item int) int {
	if debug {
		fmt.Printf("Monkey inspects an item with a worry level of %d.\n", item)
	}

	m.inspected++

	item = m.opFunc(item)

	// What does this actually do????!!!
	item = item / 3
	if debug {
		fmt.Printf("\tMonkey gets bored with item. Worry level is divided by 3 to %d.\n", item)
	}

	return item
}

func (m *Monkey) testItem(item int) bool {
	if item%m.testValue == 0 {
		if debug {
			fmt.Printf("\tCurrent worry level is divisible by %d.\n", m.testValue)
		}
		return true
	}

	if debug {
		fmt.Printf("\tCurrent worry level is not divisible by %d.\n", m.testValue)
	}
	return false
}

func (m *Monkey) throw(item int, other *Monkey) {
	if debug {
		fmt.Printf("\tItem with worry level %d is thrown to %s.\n", item, other.name)
	}
	other.catch(item)
}

func (m *Monkey) catch(item int) {
	m.items.Push(item)
}

func (m *Monkey) DoTurn(monkeys []*Monkey) {
	if debug {
		fmt.Printf("%s:\n", m.name)
	}

	for !m.items.IsEmpty() {
		item := m.items.Pop()

		item = m.inspect(item)

		if m.testItem(item) {
			m.throw(item, monkeys[m.trueMonkey])
		} else {
			m.throw(item, monkeys[m.falseMonkey])
		}
	}
}

func loadFile(input []string) []*Monkey {
	monkeys := []*Monkey{}

	var name string
	items := []int{}
	var operation string
	var testCondition string
	var resultIfTrue string
	var resultIfFalse string

	for _, s := range input {
		switch {
		case strings.HasPrefix(s, "Monkey"):
			name = strings.TrimSuffix(s, ":")
		case strings.HasPrefix(s, "Starting"):
			data := strings.SplitN(s, ":", 2)
			items = lib.StrToIntArray(data[1])
		case strings.HasPrefix(s, "Operation"):
			data := strings.SplitN(s, ":", 2)
			operation = data[1]
		case strings.HasPrefix(s, "Test"):
			data := strings.SplitN(s, ":", 2)
			testCondition = data[1]
		case strings.HasPrefix(s, "If"):
			data := strings.SplitN(s, ":", 2)

			if strings.Contains(data[0], "true") {
				resultIfTrue = data[1]
			} else {
				resultIfFalse = data[1]
			}
		default:
			m := NewMonkey(name, items, operation, testCondition, resultIfTrue, resultIfFalse)
			monkeys = append(monkeys, m)
		}
	}

	m := NewMonkey(name, items, operation, testCondition, resultIfTrue, resultIfFalse)
	monkeys = append(monkeys, m)

	fmt.Printf("monkeys loaded: %d\n", len(monkeys))

	return monkeys
}

func getTopTwoMonkeys(monkeys []*Monkey) (int, int) {
	var first, second int

	for _, m := range monkeys {
		fmt.Printf("%s inspected items %d times.\n", m.name, m.inspected)

		if m.inspected > first {
			second = first
			first = m.inspected
		} else if m.inspected > second {
			second = m.inspected
		}
	}

	return first, second
}
