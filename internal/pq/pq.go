package pq

import "container/heap"

// User Friendly interface
type PriorityQueue[T any] interface {
	Push(item T)
	Pop() *T
	IsEmpty() bool
}

// Implementation of user friendly PriorityQueue interface
type PQ[T any] struct {
	_pq[T]
	getPriority func(T) int
}

func (pq *PQ[T]) Push(item T) {
	heap.Push(&pq._pq, _item[T]{item, pq.getPriority(item)})
}
func (pq *PQ[T]) Pop() *T {
	if pq.IsEmpty() { 
		return nil
	}

	item := heap.Pop(&pq._pq).(_item[T]).value
	return &item
}
func (pq PQ[T]) IsEmpty() bool {
	return len(pq._pq) == 0
}

func NewPriorityQueue[T any](getPriority func(T) int, items ...T) PriorityQueue[T] {
	pq := PQ[T]{getPriority: getPriority}
	for _, item := range items {
		pq.Push(item)
	}
	return &pq
}
