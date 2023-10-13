package lib

import (
	"fmt"
	"strings"
)

type PriorityQueue[T any] struct {
	data []T
}

func NewPriorityQueue[T any]() *PriorityQueue[T] {
	q := PriorityQueue[T]{}

	return &q
}

func (q *PriorityQueue[T]) Push(item T) {
	q.data = append(q.data, item)
}

func (q *PriorityQueue[T]) PushAt(item T, n int) {
	if n >= len(q.data) {
		q.data = append(q.data, item)
		return
	}

	// split slice at N
	dataLeft := q.data[:n]
	dataRight := q.data[n:]

	// insert o at N
	q.data = append(dataLeft, item)

	// add the rest back in
	q.data = append(q.data, dataRight...)
}

func (q *PriorityQueue[T]) Pop() T {
	op := q.data[0]
	q.data = q.data[1:]

	return op
}

func (q *PriorityQueue[T]) PopAt(n int) (T, error) {
	if n >= len(q.data) {
		var result T
		return result, fmt.Errorf("PriorityQueue.PopN: Index out of bounds. n=%d", n)
	}

	// split slice at N
	dataLeft := q.data[:n]
	dataRight := q.data[n:]

	item := dataRight[0]

	q.data = append(dataLeft, dataRight[1:]...)

	return item, nil
}

func (q *PriorityQueue[T]) Peek() T {
	return q.data[0]
}

func (q *PriorityQueue[T]) PeekAt(n int) T {
	return q.data[n]
}

func (q *PriorityQueue[T]) IsEmpty() bool {
	return len(q.data) == 0
}

func (q *PriorityQueue[T]) Length() int {
	return len(q.data)
}

func (q *PriorityQueue[T]) String() string {
	lines := []string{}

	for i, op := range q.data {
		line := fmt.Sprintf("%d: %+v", i, op)
		lines = append(lines, line)
	}

	return strings.Join(lines, "\n")
}
