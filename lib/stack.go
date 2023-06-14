package lib

import "fmt"

type Stack[T any] []T
type StackList[T any] []Stack[T]

func (s *Stack[T]) Push(v T) {
	*s = append(*s, v)
}

func (s *Stack[T]) Pop() T {
	popped := (*s).PopN(1)

	if len(popped) > 0 {
		return popped[0]
	}

	return *new(T)
}

func (s *Stack[T]) PopN(n int) []T {
	l := len(*s)

	if l > 0 && n <= l {
		pop := (*s)[l-n:]
		*s = (*s)[:l-n]

		return pop
	}

	return nil
}

func (s *Stack[T]) String() string {
	return fmt.Sprintf("%+v\n", *s)
}
