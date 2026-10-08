func search(nums []int, target int) int {
	l:=0
	r:=len(nums)-1
	for l<=r{
		m := (l+r) /2
		
		if nums[m] == target{
			return m
		}else if nums[m] >= nums[l]{
			//左邊窗口有序時確認target 是否在其中
			if target <= nums[m] && target >= nums[l]{
				r = m-1
			}else{
				l = m+1
			}
		}else if nums[m] < nums[r]{
			//右邊窗口有序時確認target 是否在其中
			if target >= nums[m] && target <= nums[r]{
				l = m+1
			}else{
				r = m-1
			}
		}
	}
	return -1
}
