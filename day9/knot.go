package day9

import (
	"fmt"
	"math"
)

type Knot struct {
	Name string
	X    int
	Y    int
}

func (k *Knot) Move(direction string) {
	switch direction {
	case "U":
		k.Y += 1
	case "D":
		k.Y -= 1
	case "L":
		k.X -= 1
	case "R":
		k.X += 1
	}

	// fmt.Println(fmt.Sprintf("%s moved to %d, %d", k.Name, k.X, k.Y))
}

func (k *Knot) MoveMany(m Move) {
	switch m.Direction {
	case "U":
		k.Y += m.Spaces
	case "D":
		k.Y -= m.Spaces
	case "L":
		k.X -= m.Spaces
	case "R":
		k.X += m.Spaces
	}
}

func (k *Knot) Distance(o Knot) float64 {
	return math.Sqrt(math.Pow(float64(k.XDiff(o)), 2) + math.Pow(float64(k.YDiff(o)), 2))
}

func (k *Knot) XDiff(o Knot) int {
	return o.X - k.X
}

func (k *Knot) YDiff(o Knot) int {
	return o.Y - k.Y
}

func (k *Knot) IsAt(x, y int) bool {
	return k.X == x && k.Y == y
}

func (k *Knot) Follow(head Knot) {
	xDiff := k.XDiff(head)
	yDiff := k.YDiff(head)
	d := k.Distance(head)

	if math.Abs(float64(xDiff)) > 1 || d > math.Sqrt(2) {
		if xDiff < 0 {
			k.Move("L")
		} else if xDiff > 0 {
			k.Move("R")
		}
	}

	// d = k.Distance(head)

	if math.Abs(float64(yDiff)) > 1 || d > math.Sqrt(2) {
		if yDiff < 0 {
			k.Move("D")
		} else if yDiff > 0 {
			k.Move("U")
		}
	}
}

func (k *Knot) String() string {
	return fmt.Sprintf("%d-%d", k.X, k.Y)
}
