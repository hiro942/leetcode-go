package data_structures

type Queue[T any] []T

func (q *Queue[T]) Push(v T) {
	*q = append(*q, v)
}

func (q *Queue[T]) Pop() T {
	front := (*q)[0]
	*q = (*q)[1:]
	return front
}
