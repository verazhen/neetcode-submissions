func maxSlidingWindow(nums []int, k int) []int {
    decsendArr:=[][]int{} //big -> small [value,index]
	j := 0
	results:=[]int{}
	for j < len(nums) {

		//r -> l: remove ele < right element
		for i:=len(decsendArr)-1;i >= 0;i--{
			if decsendArr[i][0] > nums[j]{
				break
			}
			decsendArr = decsendArr[0:i]
		}
		//l -> r: remove expired element
		leftBound := j-k +1
		for len(decsendArr)>0 && decsendArr[0][1] < leftBound{
			decsendArr = decsendArr[1:]
		}

		//add right element
		decsendArr = append(decsendArr,[]int{nums[j],j})
		
		//cal result
		if j >= k-1{
			results = append(results,decsendArr[0][0])
		}

		j++
	}
	return results
}
