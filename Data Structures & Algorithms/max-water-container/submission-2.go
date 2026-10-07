func maxArea(heights []int) int {
	//h * w == min(heights[i],heights[j]) * (j-i)
	i:=0
	j:=len(heights)-1
	maxStorage:=0
	for i<j{
		
		maxStorage = max(min(heights[i],heights[j])*(j-i),maxStorage)

		// Move the shorter side because keeping it cannot produce a larger area
		if heights[i] <= heights[j]{
			i++
		}else{
			j--
		}
	}
	return maxStorage
}
