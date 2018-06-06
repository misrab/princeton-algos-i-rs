package medianheap

import (
	//"fmt"
	"math"

	"container/heap"
)

type median_heap struct {
	High *uint64MinHeap
	Low  *uint64MaxHeap
}

type MedianHeap interface {
	Insert(x uint64)
	ExtractMedian() uint64

	Rebalance()
}

func NewMedianHeap() MedianHeap {
	mh := new(median_heap)

	mh.High = &uint64MinHeap{}
	mh.Low = &uint64MaxHeap{}

	heap.Init(mh.High)
	heap.Init(mh.Low)

	return mh
}

func (mh *median_heap) Rebalance() {
	len_high := mh.High.Len()
	len_low := mh.Low.Len()

	diff := math.Abs(float64(len_high) - float64(len_low))
	//diff_int := diff.(int)
	if diff < 2 {
		return
	}

	//println("rebalancing")

	diff_floor := int(math.Floor(diff / 2))
	if len_high > len_low {
		for i := 0; i < diff_floor; i++ {
			heap.Push(mh.Low, heap.Pop(mh.High))
		}
	} else {
		for i := 0; i < diff_floor; i++ {
			heap.Push(mh.High, heap.Pop(mh.Low))
		}
	}

	//fmt.Printf("lens are %d and %d\n", mh.Low.Len(), mh.High.Len())

}

func (mh *median_heap) Insert(x uint64) {
	if mh.Low.Len() == 0 {
		heap.Push(mh.Low, x)
		return
	}
	if mh.High.Len() == 0 {
		heap.Push(mh.High, x)
		return
	}

	low_popped := heap.Pop(mh.Low).(uint64)
	high_popped := heap.Pop(mh.High).(uint64)

	switch {
	case x < low_popped:
		heap.Push(mh.Low, x)
		//fmt.Printf("added %d to low\n", x)
		break
	case x > high_popped:
		heap.Push(mh.High, x)
		//fmt.Printf("added %d to high\n", x)
		break
	default:
		heap.Push(mh.Low, x)
	}

	// reinsert pops
	heap.Push(mh.Low, low_popped)
	heap.Push(mh.High, high_popped)

	//fmt.Printf("lens are %d and %d\n", mh.Low.Len(), mh.High.Len())

	mh.Rebalance()
}

func (mh *median_heap) ExtractMedian() uint64 {
	total_size := mh.High.Len() + mh.Low.Len()
	var median interface{}
	// ! no empty / error checking
	// ! assuming heaps balanced

	// if its odd, median is in the odd heap
	if total_size%2 == 1 {
		if mh.High.Len()%2 == 1 {
			median = heap.Pop(mh.High)
		} else {
			median = heap.Pop(mh.Low)
		}
	} else {
		// if it's even, median is in the max heap
		median = heap.Pop(mh.Low)
	}

	mh.Rebalance()

	return median.(uint64)
}

// heaps

// An uint64MinHeap is a min-heap of uint64s.
type uint64MinHeap []uint64

func (h uint64MinHeap) Len() int           { return len(h) }
func (h uint64MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h uint64MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *uint64MinHeap) Push(x interface{}) {
	// Push and Pop use pouint64er receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(uint64))
}

func (h *uint64MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// max heap version (just change Less())

type uint64MaxHeap []uint64

func (h uint64MaxHeap) Len() int           { return len(h) }
func (h uint64MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h uint64MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *uint64MaxHeap) Push(x interface{}) {
	// Push and Pop use pouint64er receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(uint64))
}

func (h *uint64MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
