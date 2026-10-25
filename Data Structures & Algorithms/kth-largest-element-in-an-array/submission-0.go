func findKthLargest(nums []int, k int) int {
	left := 0
	right := len(nums)-1
	for{
		pivotIndex := right
		swapIndex := left
		for i:=left; i < right; i++{
			if nums[i] <= nums[pivotIndex]{
				nums[i],nums[swapIndex] = nums[swapIndex],nums[i]
				swapIndex++ // swap 最後會停在 pivot point
			}
		}
		//swap pivot point
		nums[pivotIndex],nums[swapIndex] = nums[swapIndex],nums[pivotIndex]
		if swapIndex == len(nums)-k{
			break
		}else if swapIndex < len(nums)-k{
			//find right window
			left = swapIndex + 1
		}else{
			//find left window
			right = swapIndex - 1
		}
	}
	return nums[len(nums)-k]
}
