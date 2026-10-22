type heapCustom struct{
	stones []int
} 

func lastStoneWeight(stones []int) int {
	h := heapCustom{stones:stones}
	h.heapify()
	for len(h.stones) > 1{
		bigger := h.pop()
		smaller := h.pop()
		if remain:=bigger-smaller;remain > 0{
			h.stones = append(h.stones,remain)
			h.siftUp()
		}
	}
	if len(h.stones) == 0{
		return 0
	}
	return h.stones[0]
}
func (h *heapCustom)heapify(){
	//Last Non-Leaf Node: n/2 -1
	//sift down smaller one
	for i := len(h.stones)/2 - 1; i>=0; i--{
		h.siftDown(i)
	}
}


func (h *heapCustom)pop()int{
	pop := h.stones[0]
	h.stones[0] = h.stones[len(h.stones)-1]
	h.stones = h.stones[0:len(h.stones)-1]
	h.siftDown(0)
	return pop
}
func (h *heapCustom)siftDown(parent int){
	for parent >= 0{
		left := parent*2 +1
		if left >= len(h.stones){
			break
		}
		right := parent*2 +2
		child := left
		if right < len(h.stones) && h.stones[right]>h.stones[left]{
			child = right
		} 
		if h.stones[parent] >= h.stones[child]{
			break
		}
		//swap node
		h.stones[parent], h.stones[child] = h.stones[child], h.stones[parent] 
		parent = child
	}

}


func (h *heapCustom)siftUp(){
	child := len(h.stones) - 1

	for child > 0{
		parent := (child - 1) / 2
		if h.stones[child] <= h.stones[parent]{
			break
		}
		h.stones[parent], h.stones[child] = h.stones[child], h.stones[parent] 
		child = parent
	}

}
