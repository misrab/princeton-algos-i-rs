package medianheap

import (
	"container/heap"
)

type median_heap struct {
	high uint64Heap
	low  uint64Heap
}

type MedianHeap interface {
	Insert(x uint64)
	ExtractMedian() uint64
}

func NewMedianHeap() MedianHeap {
	mh := new(median_heap)

	heap.Init(mh.high)
	heap.Init(mh.low)

	return mh
}

func (mh *median_heap) Insert(x uint64) {

}

func (mh *median_heap) ExtractMedian() uint64 {

	return 0
}

// heaps

// An uint64Heap is a min-heap of uint64s.
type uint64Heap []uint64

func (h uint64Heap) Len() int           { return len(h) }
func (h uint64Heap) Less(i, j int) bool { return h[i] < h[j] }
func (h uint64Heap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h uint64Heap) Push(x interface{}) {
	// Push and Pop use pouint64er receivers because they modify the slice's length,
	// not just its contents.
	h = append(h, x.(uint64))
}

func (h uint64Heap) Pop() interface{} {
	old := h
	n := len(old)
	x := old[n-1]
	h = old[0 : n-1]
	return x
}
