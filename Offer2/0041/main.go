package main

import "container/heap"

type MaxHeap []int
type MinHeap []int

func (h MaxHeap) Len() int            { return len(h) }
func (h MaxHeap) Less(i, j int) bool  { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MaxHeap) Pop() interface{} {
	old := *h
	lastIndex := len(old) - 1
	last := old[lastIndex]
	*h = old[:lastIndex]
	return last
}

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MinHeap) Pop() interface{} {
	old := *h
	lastIndex := len(old) - 1
	last := old[lastIndex]
	*h = old[:lastIndex]
	return last
}

type MedianFinder struct {
	MaxHeap
	MinHeap
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	m := MedianFinder{
		MaxHeap{},
		MinHeap{},
	}
	heap.Init(&m.MaxHeap)
	heap.Init(&m.MinHeap)
	return m
}

func (this *MedianFinder) AddNum(num int) {
	if this.MaxHeap.Len() == this.MinHeap.Len() {
		heap.Push(&this.MinHeap, num)
		heap.Push(&this.MaxHeap, heap.Pop(&this.MinHeap))
	} else {
		heap.Push(&this.MaxHeap, num)
		heap.Push(&this.MinHeap, heap.Pop(&this.MaxHeap))
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.MaxHeap.Len() == this.MinHeap.Len() {
		return float64(this.MaxHeap[0]+this.MinHeap[0]) / 2
	}
	return float64(this.MaxHeap[0])
}
