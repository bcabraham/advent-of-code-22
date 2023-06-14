package day9

import "fmt"

type Node struct {
	Next *Node
	Knot Knot
}

type KnotList struct {
	Head    *Node
	Size    int
	Visited map[int]map[int]bool
}

func (kl *KnotList) Add(k Knot) {
	n := Node{nil, k}

	var current *Node

	if kl.Head == nil {
		kl.Head = &n
	} else {
		current = kl.Head

		for current.Next != nil {
			current = current.Next
		}

		current.Next = &n
	}

	kl.Size += 1
}

func (kl *KnotList) IsEmpty() bool {
	return kl.Size == 0
}

func (kl *KnotList) Update(m Move) {
	current := kl.Head
	current.Knot.Move(m.Direction)

	for current.Next != nil {
		prev := current
		current = current.Next
		current.Knot.Follow(prev.Knot)
	}

	// This should be the tail after traversal
	kl.Visit(current.Knot)
}

func NewKnotList(size int) KnotList {
	visits := map[int]map[int]bool{}

	kl := KnotList{nil, 0, visits}

	for i := 0; i < size; i++ {
		var name string

		if i == 0 {
			name = "H"
		} else {
			name = fmt.Sprintf("%d", i)
		}

		k := Knot{name, 0, 0}
		kl.Add(k)
	}

	return kl
}

func (kl *KnotList) Visit(k Knot) {
	if kl.Visited[k.X] == nil {
		kl.Visited[k.X] = map[int]bool{}
	}
	kl.Visited[k.X][k.Y] = true
}

func (kl *KnotList) WasVisited(x, y int) bool {
	return kl.Visited[x] != nil && kl.Visited[x][y]
}

func (kl *KnotList) NumVisited() int {
	count := 0

	for _, v := range kl.Visited {
		for _, visited := range v {
			if visited {
				count += 1
			}
		}
	}

	return count
}

// Update board state and display
func (kl *KnotList) Show(arr Array2D) {
	// Show visited squares as *
	for x, v := range kl.Visited {
		for y, visited := range v {
			if visited {
				arr.Set(x, y, "*")
			}
		}
	}

	knots := []Knot{}
	current := kl.Head
	knots = append(knots, current.Knot)
	for current.Next != nil {
		current = current.Next
		knots = append(knots, current.Knot)
	}

	// Display Knotted squares.
	// First knots overwrite later ones if they share the same square.
	for i := len(knots) - 1; i >= 0; i-- {
		k := knots[i]
		arr.Set(k.X, k.Y, k.Name)
	}

	arr.Display()
}
