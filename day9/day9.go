package day9

import (
	"advent-of-code-22/lib"
	"fmt"
	"strings"
)

/*
Model the positions of the knots on a two-dimensional grid. Then, by following a hypothetical series of motions (your puzzle input) for the head,
you can determine how the tail will move. The head (H) and tail (T) must always be touching (diagonally adjacent and even overlapping both count as touching).

If the head is ever two steps directly up, down, left, or right from the tail, the tail must also move one step in that direction so it remains close enough.

Otherwise, if the head and tail aren't touching and aren't in the same row or column, the tail always moves one step diagonally to keep up:

You just need to work out where the tail goes as the head follows a series of motions. Assume the head and the tail both start at the same position, overlapping.

After each step, you'll need to update the position of the tail if the step means the head is no longer adjacent to the tail.

After simulating the rope, you can count up all of the positions the tail visited at least once. In this diagram, s again marks the starting position (which the tail also visited) and # marks other positions the tail visited:

Simulate your complete hypothetical series of motions. How many positions does the tail of the rope visit at least once?
*/

type Move struct {
	Direction string
	Spaces    int
}

func loadMoves(input []string) []Move {
	moves := []Move{}

	for _, data := range input {
		d := strings.SplitN(data, " ", 2)
		move := Move{d[0], lib.StrToInt(d[1])}

		moves = append(moves, move)
	}

	return moves
}

func (m *Move) String() string {
	return fmt.Sprintf("== %s %d ==", m.Direction, m.Spaces)
}

func Run() {
	input, err := lib.ReadLines("rope-bridge.txt") // rope-bridge-test.txt
	lib.HandleError(err)

	moves := loadMoves(input)
	dim := getDimensions(moves)

	fmt.Println(fmt.Sprintf("Board dimensions: %+v", dim))

	// arr.Set(center.X, center.Y, "s")

	kl := NewKnotList(10)

	fmt.Println("==", "Initial State", "==")
	kl.Show(NewArray2D(dim.Height(), dim.Width(), ".", dim.XOffset(), dim.YOffset()))

	for _, move := range moves {
		fmt.Println(move.String())
		for i := 0; i < move.Spaces; i++ {
			kl.Update(move)
		}

		// kl.Show(NewArray2D(dim.Height(), dim.Width(), ".", dim.XOffset(), dim.YOffset()))
	}

	kl.Show(NewArray2D(dim.Height(), dim.Width(), ".", dim.XOffset(), dim.YOffset()))
	fmt.Println("Spaces visited: ", kl.NumVisited())
}
