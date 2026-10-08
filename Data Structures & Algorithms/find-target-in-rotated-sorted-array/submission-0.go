func search(nums []int, target int) int {
	//1. search min index in array eg. i = 4
	l:=0
	r:=len(nums)-1
	for nums[l] > nums[r]{
		m := (l+r) /2
		if nums[m] > nums[r]{
			l = m+1
		}else{
			r = m
		}
	}

	//2. do sliding window in sorted array but get the value with rotated index
	i:=0
	j:=len(nums)-1
	for i <= j{
		k := (i+j) /2
		rk := (k+ l) % len(nums)
		if nums[rk] ==  target{
			return rk
		}else if nums[rk]  < target{
			i = k + 1
		}else{
			j = k - 1
		}
	}

	return -1
}


