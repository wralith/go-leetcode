package leetcode

import (
	"container/heap"
)

type IntHeap []int

// Implement Interface

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}
func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func lastStoneWeight(stones []int) int {
	pq := IntHeap(stones)
	heap.Init(&pq)
	for pq.Len() > 1 {
		heap.Push(&pq, heap.Pop(&pq).(int)-heap.Pop(&pq).(int))

	}
	return heap.Pop(&pq).(int)
}
