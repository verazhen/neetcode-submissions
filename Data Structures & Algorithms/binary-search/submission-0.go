func search(nums []int, target int) int {
	i:=0
	j:=len(nums)-1
	for i <= j{
		k := (i+j)/2
		if target == nums[k]{
			return k
		}else if target < nums[k]{
			j = k-1
		}else{
			i = k+1
		}
	}

	return -1
}
