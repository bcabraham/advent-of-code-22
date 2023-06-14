package day9

import (
	"fmt"
	"strings"
)

type Array2D struct {
	data    [][]string
	XOffset int
	YOffset int
}

func NewArray2D(height int, width int, defaultValue string, xOffset int, yOffset int) Array2D {
	data := make([][]string, height)

	for i := 0; i < height; i++ {
		a := make([]string, width)

		for j := range a {
			a[j] = defaultValue
		}

		data[i] = a
	}

	return Array2D{data, xOffset, yOffset}
}

func (arr Array2D) Get(x int, y int) string {
	return arr.data[y-arr.YOffset][x-arr.XOffset]
}

func (arr Array2D) Set(x int, y int, val string) {
	arr.data[y-arr.YOffset][x-arr.XOffset] = val
}

func (arr Array2D) Display() {
	output := ""
	for i := len(arr.data) - 1; i >= 0; i-- {
		output += strings.Join(arr.data[i], " ") + "\n"
	}

	fmt.Println(output)
}
