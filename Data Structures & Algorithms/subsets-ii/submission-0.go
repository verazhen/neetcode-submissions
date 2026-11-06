func subsetsWithDup(nums []int) [][]int {
	uniqueNums := []int{}
	numFreqMap := map[int]int{}
	for _,num:=range nums{
		if _,ok:=numFreqMap[num];ok{
			numFreqMap[num]++
			continue
		}
		numFreqMap[num] = 1
		uniqueNums = append(uniqueNums,num)
	}

	path:=[]int{}
	results:=[][]int{}
	var dfs func(i int)
	dfs = func(i int){
		if i >= len(uniqueNums){
			//copy result to path
			copied := append([]int{},path...)
			results = append(results,copied)
			return
		}
		
		freqCount := 0
		l := len(path)
		for freqCount <= numFreqMap[uniqueNums[i]]{
			dfs(i+1)
			path = append(path,uniqueNums[i])
			freqCount++
		}
		//initialize
		path = path[:l]
	}
	dfs(0)
	return results
}
