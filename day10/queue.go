package day10

import (
	"fmt"
	"strings"
)

type Queue struct {
	data []Operation
}

func NewQueue(input []string) Queue {
	q := Queue{[]Operation{}}

	for _, s := range input {
		op := NewOperation(s)
		q.Push(op)
	}

	return q
}

func (q *Queue) Push(o Operation) {
	q.data = append(q.data, o)
}

func (q *Queue) Pop() Operation {
	op := q.data[0]
	q.data = q.data[1:]

	return op
}

func (q *Queue) Peek() Operation {
	return q.data[0]
}

func (q *Queue) IsEmpty() bool {
	return len(q.data) == 0
}

func (q *Queue) String() string {
	lines := []string{}

	for p, op := range q.data {
		line := fmt.Sprintf("%d: %+v", p, op)
		lines = append(lines, line)
	}

	return strings.Join(lines, "\n")
}
