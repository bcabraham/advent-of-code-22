package day10

import (
	"fmt"
	"strings"
)

type PriorityQueue struct {
	data OperationList
}

func NewPriorityQueue() PriorityQueue {
	m := OperationList{}
	q := PriorityQueue{m}

	return q
}

func (q *PriorityQueue) Push(o Operation) {
	q.data = append(q.data, o)
}

func (q *PriorityQueue) PushN(n int, o Operation) {
	if n >= len(q.data) {
		q.data = append(q.data, o)
		return
	}

	// split slice at N
	dataLeft := q.data[:n]
	dataRight := q.data[n:]

	// insert o at N
	q.data = append(dataLeft, o)

	// add the rest back in
	q.data = append(q.data, dataRight...)
}

func (q *PriorityQueue) Pop() Operation {
	op := q.data[0]
	q.data = q.data[1:]

	return op
}

func (q *PriorityQueue) PopN(n int) (Operation, error) {
	if n >= len(q.data) {
		return Operation{}, fmt.Errorf("PriorityQueue.PopN: Index out of bounds. n=%d", n)
	}

	// split slice at N
	dataLeft := q.data[:n]
	dataRight := q.data[n:]

	op := dataRight[0]

	q.data = append(dataLeft, dataRight[1:]...)

	return op, nil
}

func (q *PriorityQueue) Peek() Operation {
	return q.data[0]
}

func (q *PriorityQueue) PeekN(n int) Operation {
	return q.data[n]
}

func (q *PriorityQueue) IsEmpty() bool {
	return len(q.data) == 0
}

func (q *PriorityQueue) String() string {
	lines := []string{}

	for i, op := range q.data {
		line := fmt.Sprintf("%d: %+v", i, op)
		lines = append(lines, line)
	}

	return strings.Join(lines, "\n")
}
