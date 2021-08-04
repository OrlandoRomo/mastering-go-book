package main

import (
	"container/heap"
	"fmt"
)

type heap32 []float32

func (h *heap32) Pop() interface{} {
	old := *h
	x := old[len(old)-1]
	new := old[0 : len(old)-1]
	*h = new
	return x
}

func (h *heap32) Push(x interface{}) {
	*h = append(*h, x.(float32))
}

func (h heap32) Len() int {
	return len(h)
}

func (h heap32) Less(a, b int) bool {
	return h[a] < h[b]
}

func (h heap32) Swap(a, b int) {
	h[a], h[b] = h[b], h[a]
}

func main() {
	myHeap := &heap32{1.2, 2.1, 3.1, -100.1}
	heap.Init(myHeap)
	size := myHeap.Len()
	fmt.Println("size of the heap", size)
	fmt.Println(myHeap)
	myHeap.Push(float32(45.4))
	myHeap.Push(float32(30.5))
	fmt.Println("size of the heap", myHeap.Len())
	heap.Init(myHeap)
	fmt.Println(myHeap)
}
