import stdheap "container/heap"


func lastStoneWeight(stones []int) int {
	maxHeap := MaxHeap(stones)
	stdheap.Init(&maxHeap)

	for maxHeap.Len() > 1{
		bigger := stdheap.Pop(&maxHeap).(int)
		smaller := stdheap.Pop(&maxHeap).(int)
		if remain:=bigger-smaller;remain > 0{
			stdheap.Push(&maxHeap,remain)
		}
	}
	if maxHeap.Len() == 0{
		return 0
	}
	return stdheap.Pop(&maxHeap).(int)
}



// ===== Min Heap Template =====

type MaxHeap []int

func (h MaxHeap) Len() int {
    return len(h)
}

func (h MaxHeap) Less(i, j int) bool {
    return h[i] > h[j]
}

func (h MaxHeap) Swap(i, j int) {
    h[i], h[j] = h[j], h[i]
}

func (h *MaxHeap) Push(x any) {
    *h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() any {
    old := *h
    n := len(old)

    val := old[n-1]
    *h = old[:n-1]

    return val
}
