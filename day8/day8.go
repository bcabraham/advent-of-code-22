package day8

import (
	"advent-of-code-22/lib"
	"fmt"
	"strings"
)

func Run() {
	input, err := lib.ReadLines("tree-house.txt") // tree-house-test
	lib.HandleError(err)

	arr := NewArray2D(input)
	fmt.Println(arr.Heights())
	// visibleCount := CountVisibleTrees(arr)
	// fmt.Println(arr.Visibility())
	// fmt.Println("Ttl visible:", visibleCount)

	bestScore := MaxScenicScore(arr)
	fmt.Println(arr.ScenicScores())
	fmt.Println("Best Score:", bestScore)
}

func MaxScenicScore(arr Array2D) int {
	maxScenicScore := 0
	for row := 0; row < arr.GetRowLength(); row++ {
		r := arr.GetRow(row)
		for col := 0; col < arr.GetColumnLength(row); col++ {
			c := arr.GetColumn(col)
			t := arr.Get(row, col)
			t.ScenicScore = t.CalcScenicScore(r, c)

			if t.ScenicScore > maxScenicScore {
				maxScenicScore = t.ScenicScore
			}
		}
	}

	return maxScenicScore
}

func CountVisibleTrees(arr Array2D) int {
	visibleCount := 0
	for row := 0; row < arr.GetRowLength(); row++ {
		r := arr.GetRow(row)
		for col := 0; col < arr.GetColumnLength(row); col++ {
			c := arr.GetColumn(col)
			t := arr.Get(row, col)
			t.Visible = t.IsVisible(r, c)
			// fmt.Printf("%+v\n", *t)

			if t.Visible {
				visibleCount++
			}
		}
	}

	return visibleCount
}

type Array2D [][]Tree
type Comparable func(a, b Tree) bool
type Tree struct {
	Row         int
	Column      int
	Height      int
	Visible     bool
	ScenicScore int
}

func NewArray2D(input []string) Array2D {
	arr := make(Array2D, len(input))
	for i, line := range input {
		a := make([]Tree, len(line))

		for j, h := range strings.Split(line, "") {
			a[j] = NewTree(i, j, lib.StrToInt(h))
		}

		arr[i] = a
	}

	return arr
}

func NewTree(row, col, h int) Tree {
	return Tree{row, col, h, false, 0}
}

func (arr Array2D) Heights() string {
	lines := []string{"Heights:"}

	for _, r := range arr {
		s := ""
		for _, c := range r {
			s += fmt.Sprintf("%d ", c.Height)
		}

		lines = append(lines, s)
	}

	lines = append(lines, "")

	return strings.Join(lines, "\n")
}

func (arr Array2D) Visibility() string {
	lines := []string{"Visibility:"}

	for _, r := range arr {
		s := ""
		for _, c := range r {
			if c.Visible {
				s += "T"
			} else {
				s += "F"
			}
		}

		lines = append(lines, s)
	}

	lines = append(lines, "")

	return strings.Join(lines, "\n")
}

func (arr Array2D) ScenicScores() string {
	lines := []string{"ScenicScores:"}

	for _, r := range arr {
		s := ""
		for _, c := range r {
			s += fmt.Sprintf("%-7s", fmt.Sprint(c.ScenicScore))
		}

		lines = append(lines, s)
	}

	lines = append(lines, "")

	return strings.Join(lines, "\n")
}

func (arr Array2D) Get(row, col int) *Tree {
	return &arr[row][col]
}

func (arr Array2D) GetColumnLength(row int) int {
	return len(arr[row])
}

func (arr Array2D) GetRowLength() int {
	return len(arr)
}

func (arr Array2D) GetRow(row int) []Tree {
	return arr[row]
}

func (arr Array2D) GetColumn(col int) []Tree {
	column := []Tree{}

	for _, a := range arr {
		column = append(column, a[col])
	}

	return column
}

/*
A tree is visible if all of the other trees between it and an edge
of the grid are shorter than it.

Only consider trees in the same row or column; that is,
only look up, down, left, or right from any given tree.
*/
func (t Tree) IsVisible(row, col []Tree) bool {
	// t := arr.Get(row, col)
	rowLength := len(row)
	colLength := len(col)

	// Tree is in first or last row, so visible by default
	if t.Row == 0 || t.Row == rowLength-1 {
		return true
	}
	// Tree is in first or last column, so visible by default
	if t.Column == 0 || t.Column == colLength-1 {
		return true
	}

	leftVisible, rightVisible := true, true
	for i := 1; i <= (colLength/2)+1; i++ {
		left := t.Column - i
		right := t.Column + i

		if left >= 0 {
			if row[left].Height >= t.Height {
				leftVisible = false
			}
		}
		if right < colLength {
			if row[right].Height >= t.Height {
				rightVisible = false
			}
		}
	}

	upVisible, downVisible := true, true
	for i := 1; i <= (rowLength/2)+1; i++ {
		up := t.Row - i
		down := t.Row + i

		if up >= 0 {
			if col[up].Height >= t.Height {
				upVisible = false
			}
		}
		if down < rowLength {
			if col[down].Height >= t.Height {
				downVisible = false
			}
		}
	}

	return leftVisible || rightVisible || upVisible || downVisible
}

func (t Tree) CalcScenicScore(row, col []Tree) int {
	rowLength := len(row)
	colLength := len(col)

	// If a tree is right on the edge,
	// at least one of its viewing distances will be zero
	if t.Row == 0 || t.Row == rowLength-1 || t.Column == 0 || t.Column == colLength-1 {
		return 0
	}

	leftScore, rightScore := 0, 0
	doLeft, doRight := true, true
	for i := 1; i <= (colLength)+1; i++ {
		left := t.Column - i
		right := t.Column + i

		if left >= 0 && doLeft {
			leftScore++
			if row[left].Height >= t.Height {
				doLeft = false
			}
		}

		if right < colLength && doRight {
			rightScore++
			if row[right].Height >= t.Height {
				doRight = false
			}
		}

		if !doLeft && !doRight || left < 0 && right >= colLength {
			break
		}
	}

	upScore, downScore := 0, 0
	doUp, doDown := true, true
	for i := 1; i <= (rowLength)+1; i++ {
		up := t.Row - i
		down := t.Row + i

		if up >= 0 && doUp {
			upScore++
			if col[up].Height >= t.Height {
				doUp = false
			}
		}

		if down < rowLength && doDown {
			downScore++
			if col[down].Height >= t.Height {
				doDown = false
			}
		}

		if !doUp && !doDown || up < 0 && down >= rowLength {
			break
		}
	}

	return leftScore * rightScore * upScore * downScore
}
