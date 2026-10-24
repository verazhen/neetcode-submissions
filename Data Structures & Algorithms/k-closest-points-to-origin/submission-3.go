func kClosest(points [][]int, k int) [][]int {

	//0,1,2,[7],9,10 k=3/k=4, pivotIndex=3
	pivotIndex := len(points) -1
	left:=0
	for { // left len != k
		pivot := points[pivotIndex]
		pivotDistance := pivot[0]*pivot[0] + pivot[1]*pivot[1]
		swapIndex := left
		for j := left; j < pivotIndex; j++{
			pointDistance := points[j][0]*points[j][0] + points[j][1]* points[j][1]
			if pointDistance < pivotDistance{
				//swap
				points[swapIndex],points[j] = points[j],points[swapIndex]
				swapIndex++
			}
		}
		// put pivotIndex
		points[swapIndex],points[pivotIndex] = points[pivotIndex],points[swapIndex]

		if swapIndex == k || swapIndex+1 == k{
			break
		}

		if k-1 < swapIndex{
			pivotIndex = swapIndex - 1
		}else{
			left = swapIndex + 1
		}
		
	}
	return points[0:k]
}
