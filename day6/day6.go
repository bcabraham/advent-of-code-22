package day6

import (
	"advent-of-code-22/lib"
	"fmt"
)

func Run() {
	input, err := lib.ReadLines("tuning-trouble.txt")
	lib.HandleError(err)
	signal := input[0]

	searchQueue(signal, 4)
	searchQueue(signal, 14)
}

type Queue struct {
	length int
	data   []string
}

func NewQueue(length int) Queue {
	d := make([]string, length)
	return Queue{length, d}
}

func (q *Queue) Push(s string) {
	q.data = append(q.data, s)
	if len(q.data) > q.length {
		q.data = q.data[1:]
	}
}

func (q *Queue) IsUnique() bool {
	s := set(q.data)

	// ignore case where we still have empty string
	if s[""] {
		return false
	}

	if len(q.data) == q.length && len(q.data) == len(s) {
		return true
	}

	return false
}

func set(sl []string) map[string]bool {
	var m = make(map[string]bool)

	for _, s := range sl {
		if !m[s] {
			m[s] = true
		}
	}

	return m
}

func searchQueue(signal string, length int) {
	pattern := NewQueue(length)

	for i, c := range signal {
		s := string(c)
		fmt.Println(i+1, s)

		pattern.Push(s)
		if pattern.IsUnique() {
			fmt.Printf("%+v\n", pattern.data)
			break
		}
	}
}
