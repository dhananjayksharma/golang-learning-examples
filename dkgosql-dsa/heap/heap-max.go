package main

import (
	"container/heap"
	"fmt"
)

type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] } // Max Heap
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func main() {
	nums := &MaxHeap{5, 2, 9, 1, 7}
	heap.Init(nums)

	heap.Push(nums, 3)

	fmt.Println("Max value:", (*nums)[0])

	for nums.Len() > 0 {
		fmt.Print(heap.Pop(nums), " ")
	}
}
