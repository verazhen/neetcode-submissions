func trap(height []int) int {
	i:=0
	j:=len(height)-1
	leftMax:=0
	rightMax:=0
	water := 0
	for i < j-1{
		leftMax=max(height[i],leftMax)
		rightMax=max(height[j],rightMax)
		if leftMax <= rightMax{
			water+= max(leftMax-height[i+1],0)
			i++
		}else{
			water+= max(rightMax-height[j-1],0)
			j--
		}
	}
	return water
}
