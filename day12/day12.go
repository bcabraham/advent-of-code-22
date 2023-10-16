package day12

import (
	"advent-of-code-22/lib"
	"fmt"
)

var (
	testFile    = "hill-climb-test.txt"
	testFile2   = "hill-climb-test-basic.txt"
	problemFile = "hill-climb-input.txt"
	debug       = false
)

const (
	G_COST = 50
)

/*
charValues :=
E: 69
S: 83
a: 97
z: 122
*/

func Run() {
	input, err := lib.ReadLines("day12", problemFile)
	lib.HandleError(err)

	heightmap, start, end := loadFile(input)
	fmt.Println(heightmap)

	path, nodeMap := AStar(start, end, heightmap)

	fmt.Println("Path taken:")
	for i, node := range path {
		fmt.Printf("%d: %s\n", i, node)
		heightmap[node.Row][node.Col] = '*'
	}

	fmt.Println("Path length:", len(path))
	DisplayNodeMap(nodeMap, path)
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
	Row, Col       int
	Value          rune
	Gcost, Hcost   int
	Parent         *Node
	Opened, Closed bool
}

type NodeList []*Node
type NodeMap [][]*Node

func NewNode(row, col int, val rune, parent *Node) *Node {
	return &Node{row, col, val, 0, 0, parent, false, false}
}

func (n *Node) FCost() int {
	return n.Gcost + n.Hcost
}

func (n *Node) String() string {
	return fmt.Sprintf("Row:%d, Col:%d, Value:%c, Fcost:%d, Gcost:%d, Hcost:%d", n.Row, n.Col, n.Value, n.FCost(), n.Gcost, n.Hcost)
}

func (n *Node) Equals(node *Node) bool {
	return n.Row == node.Row && n.Col == node.Col
}

func (n *Node) IsTraversable(node *Node) bool {
	return n.Value == 'S' && node.Value == 'a' || n.Value == 'z' && node.Value == 'E' || node.Value-n.Value >= -1 && node.Value-n.Value <= 1
}

func NewNodeMap(hm HeightMap) NodeMap {
	length := len(hm)
	width := len(hm[0])

	nm := make([][]*Node, length)

	for i := 0; i < length; i++ {
		nm[i] = make([]*Node, width)
	}

	for i := 0; i < length; i++ {
		for j := 0; j < width; j++ {
			nm[i][j] = NewNode(i, j, hm[i][j], nil)
		}
	}

	return nm
}

func DisplayNodeMap(nm NodeMap, nl NodeList) {

	for i := 0; i < len(nm); i++ {
		output := ""
		for j := 0; j < len(nm[0]); j++ {
			node := nm[i][j]

			if node.IsIn(nl) {
				output += fmt.Sprintf("%c*", node.Value)
			} else if node.Opened && node.Closed {
				output += fmt.Sprintf("%c-", node.Value)
			} else if node.Closed {
				output += "C "
			} else if node.Opened {
				output += "O "
			} else {
				output += fmt.Sprintf("%c ", node.Value)
			}
		}
		fmt.Println(output)
	}
}

func reconstructPath(current *Node) NodeList {
	reversed := NodeList{}

	for current.Parent != nil {
		reversed = append(reversed, current)
		current = current.Parent
	}

	nodes := NodeList{}

	for i := len(reversed) - 1; i >= 0; i-- {
		nodes = append(nodes, reversed[i])
	}

	return nodes
}

func (nm NodeMap) GetNeighbors(n *Node) NodeList {
	nodes := NodeList{}

	length := len(nm)
	width := len(nm[0])

	// left
	if n.Col > 0 {
		row := n.Row
		col := n.Col - 1
		nodes = append(nodes, nm[row][col])
	}

	// right
	if n.Col < width-1 {
		row := n.Row
		col := n.Col + 1
		nodes = append(nodes, nm[row][col])
	}

	// up
	if n.Row > 0 {
		row := n.Row - 1
		col := n.Col
		nodes = append(nodes, nm[row][col])
	}

	// down
	if n.Row < length-1 {
		row := n.Row + 1
		col := n.Col
		nodes = append(nodes, nm[row][col])
	}

	return nodes
}

func (n *Node) IsIn(nodes NodeList) bool {
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

func GetLowestFCost(nl NodeList) (NodeList, *Node) {
	nodes := NodeList{}

	if len(nl) < 1 {
		return nodes, nil
	}

	if len(nl) < 2 {
		return nodes, nl[0]
	}

	current := nl[0]
	for i := 1; i < len(nl); i++ {
		n := nl[i]
		if n.FCost() < current.FCost() || n.FCost() == current.FCost() && n.Hcost < current.Hcost {
			nodes = append(nodes, current)
			current = n
		} else {
			nodes = append(nodes, n)
		}
	}

	return nodes, current
}

func AStar(start, end Index, hm HeightMap) (NodeList, NodeMap) {
	startNode := NewNode(start.Y, start.X, 'S', nil)
	endNode := NewNode(end.Y, end.X, 'E', nil)

	nodeMap := NewNodeMap(hm)

	openSet := NodeList{startNode}
	startNode.Opened = true

	closedSet := NodeList{}

	var current *Node

	for len(openSet) > 0 {
		debugPrint(fmt.Sprint("OpenSet length:", len(openSet)))

		openSet, current = GetLowestFCost(openSet)
		closedSet = append(closedSet, current)
		current.Closed = true

		debugPrint(fmt.Sprint("current:", current))

		if current.Equals(endNode) {
			debugPrint("current == end ? true")
			break
		}

		debugPrint("current == end ? false")

		neighbors := nodeMap.GetNeighbors(current)
		debugPrint(fmt.Sprint("Neighbors found:", len(neighbors)))

		for _, neighbor := range neighbors {
			debugPrint(fmt.Sprint("neighbor:", neighbor))
			debugPrint(fmt.Sprintf("neighbor is traversable: %t\n", current.IsTraversable(neighbor)))
			debugPrint(fmt.Sprintf("neighbor is closed: %t\n", neighbor.IsIn(closedSet)))

			if !current.IsTraversable(neighbor) || neighbor.IsIn(closedSet) {
				debugPrint("skipping.")
				continue
			}

			// gCost: how far away node is from starting node
			newGcost := current.Gcost + G_COST

			if newGcost < neighbor.Gcost || !neighbor.IsIn(openSet) {
				neighbor.Gcost = newGcost

				// hCost: how far away node is from end node
				neighbor.Hcost = neighbor.GetDistance(end)
				neighbor.Parent = current

				debugPrint(fmt.Sprint("neighbor cost updated:", neighbor))

				if !neighbor.IsIn(openSet) {
					openSet = append(openSet, neighbor)
					debugPrint("neighbor added to openSet")
					neighbor.Opened = true
				}
			}

		}
	}

	return reconstructPath(current), nodeMap
}

func debugPrint(str string) {
	if debug {
		fmt.Println(str)
	}
}
