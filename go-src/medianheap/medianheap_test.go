package medianheap

import (
	//"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBasic(t *testing.T) {
	mh := NewMedianHeap()
	mh_struct := mh.(*median_heap)

	mh.Insert(1)
	mh.Insert(2)

	len_low := mh_struct.Low.Len()
	len_high := mh_struct.High.Len()

	assert.True(t, len_low == len_high && len_low == 1, "lens should be 1")

	mh.Insert(4)
	mh.Insert(3)

	len_low = mh_struct.Low.Len()
	len_high = mh_struct.High.Len()
	assert.True(t, len_low == len_high && len_low == 2, "lens should be 2")
	//fmt.Printf("lens %d and %d\n%+v", len_low, len_high, mh_struct)

	median := mh.ExtractMedian()
	assert.Equal(t, median, uint64(3))
}
