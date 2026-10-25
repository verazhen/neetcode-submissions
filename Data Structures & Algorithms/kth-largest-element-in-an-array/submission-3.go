//3-way partition
// [left]---[lt]--[i]--[gt]---[right]
// nums[i] < pivot	Swap(i, lt)、lt++、i++	擴大小於區間
// nums[i] == pivot	i++	擴大等於區間
// nums[i] > pivot	Swap(i, gt)、gt--	交換過來的元素尚未檢查
func findKthLargest(nums []int, k int) int {
	left := 0
	right := len(nums)-1
	for{
		pivot := nums[right]
		lt := left
		gt := right 
		i := left
		for i <= gt{
			if nums[i] < pivot{
				nums[i],nums[lt] = nums[lt],nums[i]
				lt++
				i++
			}else if nums[i] == pivot{
				i++
			}else{
				nums[i],nums[gt] = nums[gt],nums[i]
				gt--
				//i停留在原位，檢查gt swap過來的值
			}
		}

		if len(nums)-k >= lt && len(nums)-k <= gt{
			break
		}else if len(nums)-k > gt{
			//find right window
			left = gt + 1
		}else{
			//find left window
			right = lt - 1
		}
	}
	return nums[len(nums)-k]
}

