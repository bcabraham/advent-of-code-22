package day12

import (
	"advent-of-code-22/lib"
	"fmt"
)

var (
	testFile    = "hill-climb-test.txt"
	problemFile = "hill-climb-input.txt"
)

type HeightMap [][]string

func (h HeightMap) String() string {
	var output string
	for _, line := range h {
		for _, height := range line {
			output += fmt.Sprintf("%s ", height)
		}
		output += "\n"
	}

	return output
}

func Run() {
	input, err := lib.ReadLines("day12", testFile)
	lib.HandleError(err)

	heightmap := loadFile(input)
	fmt.Println(heightmap)
}

func newHeightMap(length, width int) HeightMap {
	hm := make([][]string, length)

	for i := 0; i < length; i++ {
		hm[i] = make([]string, width)
	}

	return hm
}

func loadFile(input []string) HeightMap {
	heightmap := newHeightMap(len(input), len(input[0]))

	for i, line := range input {
		for j, height := range line {
			heightmap[i][j] = string(height)
		}
	}

	return heightmap
}
