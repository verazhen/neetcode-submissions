func productExceptSelf(nums []int) []int {
	
	numLeft:=1
	results:=[]int{1}
	numRight:=1
	right:=[]int{1}
	//cal left & right products
	for i:=0;i<len(nums)-1;i++{
		numLeft = numLeft *nums[i]
		results = append(results, numLeft)
		numRight = numRight *nums[len(nums)-i-1]
		right = append(right, numRight)
	}
	//l x r
	for i:=range results{
		results[i] = results[i] * right[len(right)-i-1] 
	}
	return results
}
