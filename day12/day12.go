package day12

import (
	"advent-of-code-22/lib"
	"fmt"
)

var (
	testFile    = "hill-climb-test.txt"
	testFile2   = "hill-climb-test-basic.txt"
	problemFile = "hill-climb-input.txt"
)

func Run() {
	input, err := lib.ReadLines("day12", testFile)
	lib.HandleError(err)

	heightmap, start, end := loadFile(input)
	fmt.Println(heightmap)

	path := AStar(start, end, heightmap)

	fmt.Println("Path length:", len(path))
}

type Index struct {
	Y, X int
}

type HeightMap [][]rune

func (h HeightMap) String() string {
	var output string
	for _, line := range h {
		for _, height := range line {
			output += fmt.Sprintf("%c ", height)
		}
		output += "\n"
	}

	return output
}

func newHeightMap(length, width int) HeightMap {
	hm := make([][]rune, length)

	for i := 0; i < length; i++ {
		hm[i] = make([]rune, width)
	}

	return hm
}

func loadFile(input []string) (HeightMap, Index, Index) {
	hm := newHeightMap(len(input), len(input[0]))
	var start Index
	var end Index

	for i, line := range input {
		for j, height := range line {
			hm[i][j] = height

			if height == 'S' {
				start = Index{i, j}
			}

			if height == 'E' {
				end = Index{i, j}
			}

		}
	}

	return hm, start, end
}

type Node struct {
	Row, Col            int
	Value               rune
	Fcost, Gcost, Hcost int
	Parent              *Node
}

func NewNode(row, col int, val rune, parent *Node) *Node {
	return &Node{row, col, val, 0, 0, 0, parent}
}

func (n *Node) Equals(node *Node) bool {
	return n.Row == node.Row && n.Col == node.Col
}

func (n *Node) IsTraversable(node *Node) bool {
	return n.Value == 'S' && node.Value == 'a' || n.Value == 'z' && node.Value == 'E' || node.Value-n.Value == 1 || node.Value-n.Value == 0
}

func reconstructPath(current *Node) []*Node {
	reversed := []*Node{}

	for current.Parent != nil {
		reversed = append(reversed, current)
		current = current.Parent
	}

	nodes := []*Node{}

	for i := len(reversed) - 1; i >= 0; i-- {
		nodes = append(nodes, reversed[i])
	}

	return nodes
}

func (n *Node) GetNeighbors(hm HeightMap) []*Node {
	nodes := []*Node{}

	length := len(hm)
	width := len(hm[0])

	// left
	if n.Col > 0 {
		row := n.Row
		col := n.Col - 1
		nodes = append(nodes, NewNode(row, col, hm[row][col], n))
	}

	// right
	if n.Col < width-1 {
		row := n.Row
		col := n.Col + 1
		nodes = append(nodes, NewNode(row, col, hm[row][col], n))
	}

	// up
	if n.Row > 0 {
		row := n.Row - 1
		col := n.Col
		nodes = append(nodes, NewNode(row, col, hm[row][col], n))
	}

	// down
	if n.Row < length-1 {
		row := n.Row + 1
		col := n.Col
		nodes = append(nodes, NewNode(row, col, hm[row][col], n))
	}

	return nodes
}

func (n *Node) IsIn(nodes []*Node) bool {
	for _, node := range nodes {
		if n.Col == node.Col && n.Row == node.Row {
			return true
		}
	}

	return false
}

func (n *Node) GetDistance(index Index) int {
	return (index.Y-n.Row)*(index.Y-n.Row) + (index.X-n.Col)*(index.X-n.Col)
}

func AStar(start, end Index, hm HeightMap) []*Node {
	startNode := NewNode(start.Y, start.X, 'S', nil)
	endNode := NewNode(end.Y, end.X, 'E', nil)
	openSet := lib.NewPriorityQueue[*Node]()
	openSet.Push(startNode)

	closedSet := []*Node{}

	var current *Node

	for !openSet.IsEmpty() {
		current = openSet.Pop()
		closedSet = append(closedSet, current)

		if current.Equals(endNode) {
			return reconstructPath(current)
		}

		for _, neighbor := range current.GetNeighbors(hm) {
			if !current.IsTraversable(neighbor) || neighbor.IsIn(closedSet) {
				continue
			}

			// gCost: how far away node is from starting node
			neighbor.Gcost = current.Gcost + 1

			// hCost: how far away node is from end node
			neighbor.Hcost = neighbor.GetDistance(end)

			// fCost: sum of gCost + hCost
			neighbor.Fcost = neighbor.Gcost + neighbor.Hcost

			neighbor.Parent = current

			skip := false
			for i := 0; i < openSet.Length(); i++ {
				if neighbor.Equals(openSet.PeekAt(i)) {
					skip = true
					break
				}
			}

			if skip {
				continue
			}

			inserted := false
			for i := 0; i < openSet.Length(); i++ {
				openNode := openSet.PeekAt(i)
				if openNode.Fcost > neighbor.Fcost || openNode.Fcost == neighbor.Fcost && openNode.Hcost > neighbor.Hcost {
					openSet.PushAt(neighbor, i)
					inserted = true
					break
				}
			}

			if !inserted {
				openSet.Push(neighbor)
			}
		}
	}

	return []*Node{}
}
