func twoSum(nums []int, target int) []int {
    complementMap := map[int]int{}//num to index map
	for i,num:=range nums{
		complement := target-num
		if j,ok:=complementMap[complement];ok{
return []int{j,i}
		}
		complementMap[num] = i
	}
	return []int{}
}
