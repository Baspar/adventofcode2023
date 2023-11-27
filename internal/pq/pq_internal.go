package pq

// PriorityQueue underlying item
type _item[T any] struct {
	value    T
	priority int
}

// Wrapper for container/heap Interface
type _pq[T any] []_item[T]

func (pq _pq[T]) Len() int { return len(pq) }

func (pq _pq[T]) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq _pq[T]) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *_pq[T]) Push(x any) {
	*pq = append(*pq, x.(_item[T]))
}

func (pq *_pq[T]) Pop() any {
	old := *pq
	item := old[len(old)-1]
	*pq = old[:len(old)-1]
	return item
}
