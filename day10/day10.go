package day10

import (
	"advent-of-code-22/lib"
	"fmt"
)

func Run() {
	input, err := lib.ReadLines("day10", "cathode-ray-tube.txt") // cathode-ray-tube-test.txt
	lib.HandleError(err)

	processor := NewProcessor(input, false)

	for processor.Cycle() {
	}

	fmt.Printf("Signal Strengths: %v\tTotal: %d\n", processor.GetSignals(), processor.SumSignals())
}
