package day10

import (
	"advent-of-code-22/lib"
	"fmt"
	"strings"
)

type Operation struct {
	Cmd string
	Arg int
}

type OperationList []Operation

func NewOperation(str string) Operation {
	data := strings.Split(str, " ")
	if len(data) < 1 {
		panic("NewOperation: Not enough args")
	}

	op := Operation{}
	op.Cmd = data[0]
	if len(data) > 1 {
		op.Arg = lib.StrToInt(data[1])
	}

	return op
}

func (o *Operation) String() string {
	return fmt.Sprintf("Cmd: %6s Arg: %5d", o.Cmd, o.Arg)
}
