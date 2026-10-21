
import stdheap "container/heap"

// ===== Min Heap Template =====

type MinHeap []int

func (h MinHeap) Len() int {
    return len(h)
}

func (h MinHeap) Less(i, j int) bool {
    return h[i] < h[j]
}

func (h MinHeap) Swap(i, j int) {
    h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x any) {
    *h = append(*h, x.(int))
}

func (h *MinHeap) Pop() any {
    old := *h
    n := len(old)

    val := old[n-1]
    *h = old[:n-1]

    return val
}

// ===== LeetCode 703 =====

type KthLargest struct {
    minHeap *MinHeap
    k       int
}

func Constructor(k int, nums []int) KthLargest {
    kl := KthLargest{
        minHeap: &MinHeap{},
        k:       k,
    }

    stdheap.Init(kl.minHeap)

    for _, num := range nums {
        kl.Add(num)
    }

    return kl
}

func (kl *KthLargest) Add(val int) int {
    if kl.minHeap.Len() < kl.k {
        stdheap.Push(kl.minHeap, val)
    } else if val > (*kl.minHeap)[0] {
        stdheap.Pop(kl.minHeap)
        stdheap.Push(kl.minHeap, val)
    }

    return (*kl.minHeap)[0]
}
