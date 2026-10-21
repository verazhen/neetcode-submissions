type KthLargest struct {
    heap []int
	k int
}


func Constructor(k int, nums []int) KthLargest {
    heap:= []int{}
	for _,num:=range nums{
		heap = addAndSwap(heap,num,k)	
	}
	return KthLargest{heap:heap,k:k}
}


func (this *KthLargest) Add(val int) int {
    this.heap = addAndSwap(this.heap,val,this.k)
	return this.heap[0]
}

func addAndSwap(heap []int,num int,k int) []int{
	//append element
	if len(heap) < k{
		heap = append(heap,num)
		return siftUp(heap)
	}else if root := heap[0]; num > root{
		heap[0] = num
		return siftDown(heap)
	}else{
		return heap
	}
	
}

func siftUp(heap []int)[]int{
	i:=len(heap) -1
	//當 i = 0，代表已經到達 Root。不用再比較
	for i > 0{
		child := i
		parent := (i-1)/2
		if heap[parent] <= heap[child]{
			break
		}
		heap[parent], heap[child] = heap[child], heap[parent]
		i = parent
	}
	return heap
}

func siftDown(heap []int)[]int{
	i:=0
	for i < len(heap){
		root := i
		if 2*i+1 >= len(heap){
			break
		}
		child := 2*i+1
		if 2*i+2 < len(heap) && heap[2*i+2] < heap[2*i+1]{
			child = 2*i+2
		}
		if heap[root] <= heap[child]{
			break
		}
		heap[root], heap[child] = heap[child], heap[root]
		i = child
	}
	return heap
}

