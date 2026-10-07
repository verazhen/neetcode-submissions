func longestConsecutive(nums []int) int {
	//stores exists number in hashmap
	numMap := map[int]struct{}{}
	for _, num:= range nums{
		numMap[num] = struct{}{}
	}
	//fo number in the array, we'll check if it is the first element in a consecutive sequence, if so, calculate the length and check if its the max so far
	maxLen:=0
	for _, num:= range nums{
		if _,ok :=numMap[num-1];ok{
			continue
		}
		l:=1
		for {
			num++
			if _,ok :=numMap[num];!ok{
				break
			}
			l++
		}
		maxLen = max(maxLen,l)
	}
	return maxLen
}
