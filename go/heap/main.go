package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
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

func main() {
	fmt.Println("vim-go")
	h := &IntHeap{2, 1, 5}
	heap.Init(h)
	fmt.Printf("h:%v\n", h)
	heap.Push(h, 3)
	fmt.Printf("h:%v\n", h)

	for h.Len() > 0 {
		fmt.Printf("%d\n", heap.Pop(h))
	}
	heap.Pop(h)
}
