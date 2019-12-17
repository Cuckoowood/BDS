// heap
package BDS

// 大顶堆int
type IntHeapM []int

// 小顶堆int
type IntHeap []int

// 专门为矩阵设计的多路归并思想的堆结构。
type IntHeapSpec [][3]int

type IntHeapSpecM [][3]int

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

func (h IntHeapSpec) Len() int           { return len(h) }
func (h IntHeapSpec) Less(i, j int) bool { return h[i][2] < h[j][2] }
func (h IntHeapSpec) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeapSpec) Push(x interface{}) {
	result := x.([3]int)
	j, y, value := result[0], result[1], result[2]
	f := [3]int{
		j, y, value,
	}
	*h = append(*h, f)
}
func (h *IntHeapSpec) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func (h *IntHeapSpec) Top() interface{} {

	return (*h)[0]
}


// 实现大顶堆的方式。
func (h IntHeapSpecM) Len() int           { return len(h) }
func (h IntHeapSpecM) Less(i, j int) bool { return h[i][2] > h[j][2] }
func (h IntHeapSpecM) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeapSpecM) Push(x interface{}) {
	result := x.([3]int)
	j, y, value := result[0], result[1], result[2]
	f := [3]int{
		j, y, value,
	}
	*h = append(*h, f)
}
func (h *IntHeapSpecM) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func (h *IntHeapSpecM) Top() interface{} {

	return (*h)[0]
}