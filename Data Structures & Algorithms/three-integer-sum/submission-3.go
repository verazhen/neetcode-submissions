func threeSum(nums []int) [][]int {
	results := [][]int{}
	if len(nums) == 0{
		return results
	}
	sort.Ints(nums)

	target := nums[0]
	for i:=0;i<len(nums);i++{
		if nums[i] >0 {
			break
		}
		if i > 0 && nums[i] == target{
			continue
		}
		target = nums[i]
		l := i + 1
		r := len(nums) - 1
		for l < r{

			if l > i+1 && nums[l] == nums[l-1]{
				l++
				continue
			}else if r < len(nums)-1 && nums[r] == nums[r+1]{
				r--
				continue
			}

			if nums[l] + nums[r] == target*-1{
				results = append(results, []int{nums[i],nums[l],nums[r]})
				//move l
				l++
				r--
			}else if nums[l] + nums[r] < target*-1{
				//move l
				l++
			}else{
				//move r
				r--
			}
		}
	}
	return results
}
