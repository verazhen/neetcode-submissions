import stdheap "container/heap"

//k 很大時適用：
// init: On 
// pop: Onlogn
// space: On
func kClosest(points [][]int, k int) [][]int {
	maxHeap := MaxHeap{}
	for _,point:=range points{
		//cal O(n)
		x:=point[0]
		y:=point[1]
		maxHeap = append(maxHeap, pointInfo{
			point:point,
			distance:x*x+y*y,})
	}
	//init: O(n)
	stdheap.Init(&maxHeap)
	//pop: (n-k)O(logn)
	for maxHeap.Len() >k{
		stdheap.Pop(&maxHeap)
	}
	results:=[][]int{}
	for _, pointInfo:= range maxHeap{
		results = append(results, pointInfo.point)
	}
	return results
}



// ===== Max Heap Template =====
type pointInfo struct{
	point []int
	distance int
}
type MaxHeap []pointInfo

func (h MaxHeap) Len() int {
    return len(h)
}

func (h MaxHeap) Less(i, j int) bool {
    return h[i].distance > h[j].distance
}

func (h MaxHeap) Swap(i, j int) {
    h[i], h[j] = h[j], h[i]
}

func (h *MaxHeap) Push(x any) {
    *h = append(*h, x.(pointInfo))
}

func (h *MaxHeap) Pop() any {
    old := *h
    n := len(old)

    val := old[n-1]
    *h = old[:n-1]

    return val
}

