package BDS

import (
	"fmt"
	"sync"
)

type Heap []int

// return a new Heap
func NewHeap() Heap {
	h := make([]int, 0)
	return h
}

func (h Heap) Len() int           { return len(h) }
func (h Heap) Less(i, j int) bool { return h[i] < h[j] }
func (h Heap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *Heap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// ---
// vaue []interface{}
// only allow:
//  uint uint8 uint16 uint32 uint64 int int8 int16 int32 int64
// float32 float64 byte rune
type HeapMe struct {
	value []interface{}
	count int // count how many data we hava used.
	cap   int
	sync  sync.Mutex
}

// 往堆中加入数据，并堆化
func (h *HeapMe) Pop() interface{} {
	return nil
}

// 往堆中拔出来数据并堆化
func (h *HeapMe) Push(data interface{}) error {
	if !t(data) {
		return fmt.Errorf("输入的类型不被支持，仅支持数字类型")
	}

	return nil

}

// 遍历这个堆中的数据，并且可以提供多种遍历方法，非递归方式。
func (h *HeapMe) Range() []interface{} {
	return nil
}

// 判断进入的信息是什么类型,
func t(data interface{}) bool {
	switch data.(type) {
	case uint, uint8, uint16, uint32, uint64, int, int8, int16, int64, float32, float64:
		return true
	default:
		return false
	}
}

// 从下到上的堆化
func heapDownToTop([]interface{}) {

}

// 从上到下的堆化
func heapTopToDown([]interface{}) {

}
