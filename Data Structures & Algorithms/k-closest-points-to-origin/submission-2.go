//k 很小時適用：
// init: O(nlogk)
// space: O(k)
func kClosest(points [][]int, k int) [][]int {
	maxHeap := MaxHeap{}
	for _,point:=range points{
		//cal O(n)
		x:=point[0]
		y:=point[1]
		distance:=x*x+y*y
		info := pointInfo{
			point: point,
			distance: distance,
		}

		if len(maxHeap) < k{
			maxHeap.Add(info)
		}else if distance < maxHeap[0].distance{
			maxHeap.Swap(info)
		}
	}

	results:=[][]int{}
	for _, pointInfo:= range maxHeap{
		results = append(results, pointInfo.point)
	}
	return results
}
func (h *MaxHeap) Add(p pointInfo){
	//add and sift up
	*h = append(*h,p)
	child := len(*h)-1
	for child > 0{
		parent := (child-1) /2
		if (*h)[child].distance <= (*h)[parent].distance {
			break
		}
		(*h)[child], (*h)[parent] = (*h)[parent], (*h)[child]
		child = parent
	}

}
func (h *MaxHeap) Swap(p pointInfo){
	//swap and sift down
	(*h)[0] = p
	parent := 0
	for parent < len(*h) {
		child := parent*2 +1
		if child >= len(*h){
			break
		}
		if child+1 < len(*h) && (*h)[child].distance < (*h)[child+1].distance{
			child = child+1
		}
		if (*h)[child].distance <= (*h)[parent].distance {
			break
		}
		(*h)[child], (*h)[parent] = (*h)[parent], (*h)[child]
		parent = child
	}
}

type pointInfo struct{
	point []int
	distance int
}
type MaxHeap []pointInfo
