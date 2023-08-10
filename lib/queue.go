package lib

import (
	"fmt"
)

type Queue[T any] struct {
	data []T
}

func NewQueue[T any](input []T) *Queue[T] {
	q := Queue[T]{}

	for _, i := range input {
		q.Push(i)
	}

	return &q
}

func (q *Queue[T]) Push(item T) {
	q.data = append(q.data, item)
}

func (q *Queue[T]) Pop() T {
	op := q.data[0]
	q.data = q.data[1:]

	return op
}

func (q *Queue[T]) Peek() T {
	return q.data[0]
}

func (q *Queue[T]) IsEmpty() bool {
	return len(q.data) == 0
}

func (q *Queue[T]) Len() int {
	return len(q.data)
}

func (q *Queue[T]) String() string {
	// lines := []string{}

	// for i, item := range q.data {
	// 	line := fmt.Sprintf("%d: %+v", i, item)
	// 	lines = append(lines, line)
	// }

	// return strings.Join(lines, "\n")
	return fmt.Sprintf("%v", q.data)
}
