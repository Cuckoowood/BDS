package BDS

// 大顶堆
type IntHeapM []int

// 小顶堆
type IntHeap []int

func (h IntHeapM) Len() int           { return len(h) }
func (h IntHeapM) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeapM) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeapM) Push(x interface{}) {

	*h = append(*h, x.(int))
}
func (h *IntHeapM) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func (h *IntHeapM) Top() interface{} {

	return (*h)[0]
}

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) {

	*h = append(*h, x.(int))
}
func (h *IntHeap) Pop() interface{} {
	old := *h
	//n := len(old)
	x := old[0]
	*h = old[1:]
	return x
}
func (h *IntHeap) Top() interface{} {

	return (*h)[0]
}
