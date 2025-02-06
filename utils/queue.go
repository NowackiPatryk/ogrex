package utils

import (
	"fmt"
	"slices"
)

type FifoQueue[T any] struct {
	values []T
}

func NewFifoQueue[T any](initial ...T) *FifoQueue[T] {
	return &FifoQueue[T]{values: initial}
}

func (q *FifoQueue[T]) Insert(value T) {
	q.values = slices.Insert(q.values, 0, value)
}

func (q *FifoQueue[T]) TakeLast() (T, error) {
	queueLength := len(q.values)

	if queueLength == 0 {
		var nullVal T
		return nullVal, fmt.Errorf("queue is empty")
	}

	lastElem := q.values[queueLength-1]
	q.values = q.values[:queueLength-1]

	return lastElem, nil
}

func (q *FifoQueue[T]) Peek() (T, error) {
	queueLength := len(q.values)

	if queueLength == 0 {
		var nullVal T
		return nullVal, fmt.Errorf("queue is empty")
	}

	lastElem := q.values[queueLength-1]
	return lastElem, nil
}

func (q *FifoQueue[T]) GetLength() int {
	return len(q.values)
}
