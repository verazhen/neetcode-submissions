func topKFrequent(nums []int, k int) []int {
	freqMap := map[int]int{} //num to freq map
	for _,num:=range nums{
		freqMap[num]++
	}

	freqArr := [10001][]int{} //freq to nums array
	for num,freq:=range freqMap{
		freqArr[freq] = append(freqArr[freq], num)
	}

	result:=[]int{}
	i:= 10000
	for {
		if i==0{
			return result
		}

		for _,num:=range freqArr[i]{
			result = append(result,num)
			if len(result)==k{
				return result
			}
		}

		i--

	}

	return result
}
