package main

import (
    "container/heap"
    "fmt"
)

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] } // Min Heap
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x any) {
    *h = append(*h, x.(int))
}

func (h *MinHeap) Pop() any {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[:n-1]
    return x
}

func main() {
    nums := &MinHeap{5, 2, 9, 1, 7}
    heap.Init(nums)

    heap.Push(nums, 3)

    fmt.Println("Min value:", (*nums)[0])

    for nums.Len() > 0 {
        fmt.Print(heap.Pop(nums), " ")
    }
}