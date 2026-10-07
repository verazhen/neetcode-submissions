func trap(height []int) int {
	//由後往前找到當前的最高值
	//[3,3,3,3,3,3,3,3,2,1]
	//[3,2,1,5] -> [3,5,5,5]
	//[3,2,4,5] -> [3,5,5,5]
	maxRight:=0
	reverseMaxArr:=make([]int, len(height))
	for j:=len(height)-1;j>=0;j--{
		maxRight = max(maxRight,height[j])
		reverseMaxArr[j] = maxRight
	}

	maxLeft:=height[0]
	i:=1
	water:=0
	for i+1 < len(height){
		maxRight = reverseMaxArr[i+1]
		water += max(min(maxRight,maxLeft) - height[i],0)
		maxLeft = max(maxLeft,height[i])
		i++
	}


	return water
}
