func largestRectangleArea(heights []int) int {

	stack := [][2]int{} //height to index
	maxArea := 0

	j:=0
	for j < len(heights){
		//if larger than the top element of stack
		//resolve the qualifying top elements
		h:=heights[j]
		startIndex := j
		for last:=len(stack)-1;last>=0;last-- {
			if h > stack[last][0]{
				break
			}
			startIndex = stack[last][1]
			maxArea = max(maxArea,stack[last][0] * (j-startIndex))
			stack = stack[:last]
		}
		//important: record the last pop index, if no pop, record j
		stack = append(stack,[2]int{h,startIndex})
		j++
	}

	//resolve all the remaining elements
	j=len(heights)-1
	last:=len(stack)-1
	for last>=0 {
		maxArea = max(maxArea,stack[last][0] * (j-stack[last][1]+1))
		stack = stack[:last]

		last--
	}

	return maxArea
}
// 1,2,3,1,2,3,1
// 1,2,3 <-1 
//