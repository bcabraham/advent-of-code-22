package day5

import (
	"advent-of-code-22/lib"
	"fmt"
	"strings"
)

type Stack []string
type StackList []Stack

func (s *Stack) Push(v string) {
	*s = append(*s, v)
}

func (s *Stack) Pop(n int) []string {
	l := len(*s)

	if l > 0 && n <= l {
		pop := (*s)[l-n:]
		*s = (*s)[:l-n]

		return pop
	}

	return nil
}

func (s *Stack) String() string {
	return fmt.Sprintf("%+v\n", *s)
}

func (sl *StackList) show() {
	fmt.Println(strings.Repeat("-", 20))
	for i, s := range *sl {
		fmt.Printf("%d: %s\n", i+1, s)
	}
	fmt.Println(strings.Repeat("-", 20))
}

type Instruction struct {
	move int
	from int
	to   int
}

func (i *Instruction) String() string {
	return fmt.Sprintf("move %d from %d to %d", (*i).move, (*i).from, (*i).to)
}

type InstructionList []Instruction

func Run() {
	input, err := lib.ReadLines("day5", "supply-stacks.txt")
	lib.HandleError(err)

	stacks, instructions := parseInput(input)

	processInstructions(stacks, instructions)
}

func processInstructions(stacks StackList, ins InstructionList) {
	stacks.show()

	for _, i := range ins {
		fmt.Println(i.String())
		stacks = processInstruction(stacks, i)
		stacks.show()
	}

	stacks.show()
}

func processInstruction(stacks StackList, ins Instruction) StackList {
	crate := stacks[ins.from-1].Pop(ins.move)

	for _, c := range crate {
		stacks[ins.to-1].Push(c)
	}

	return stacks
}

// func processInstruction(stacks StackList, ins Instruction) StackList {
// 	for i := 0; i < ins.move; i++ {
// 		crate := stacks[ins.from-1].Pop(1)
// 		stacks[ins.to-1].Push(crate)
// 	}

// 	return stacks
// }

func parseInput(input []string) (StackList, InstructionList) {
	stackInput := []string{}
	instructionsInput := []string{}

	moreStacks := true
	for _, s := range input {
		if len(s) == 0 {
			moreStacks = false
			continue
		}

		if moreStacks {
			stackInput = append(stackInput, s)
		} else {
			instructionsInput = append(instructionsInput, s)
		}
	}

	return loadStacks(stackInput), loadInstructions(instructionsInput)
}

func loadStacks(input []string) StackList {
	stacks := StackList{}

	for _, s := range input {
		crates := strings.Split(s, " ")
		stack := Stack{}
		for _, c := range crates {
			stack.Push(c)
		}

		stacks = append(stacks, stack)
	}

	return stacks
}

func loadInstructions(input []string) InstructionList {
	instructions := InstructionList{}

	for _, s := range input {
		moves := strings.Split(s, " ")
		move, from, to := lib.StrToInt(moves[1]), lib.StrToInt(moves[3]), lib.StrToInt(moves[5])
		instruction := Instruction{move, from, to}
		instructions = append(instructions, instruction)
	}

	return instructions
}
