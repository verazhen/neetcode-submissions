func findMin(nums []int) int {
	// 1. l < r : return l
	// 2. l > r : 
	// -- a. m > r : l to m+1
	// -- b. m < r : r to m
	l := 0
	r := len(nums) - 1
	for nums[l] > nums[r]{
		m := (l + r) /2
		if nums[m] > nums[r]{
			l = m+1
		}else{
			r = m
		}
	}

	return nums[l]
}
