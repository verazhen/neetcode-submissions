func combinationSum(nums []int, target int) [][]int {
    path:=[]int{}
	results:=[][]int{}
	var dfs func(i int,sum int)
	dfs = func(i int,sum int){
		if i >= len(nums){
			return
		}

		newSum := sum
		l := len(path)
		for newSum < target{
			dfs(i+1,newSum)
			path = append(path,nums[i])
			newSum += nums[i]
		}
		if newSum == target{			
			copied := append([]int{},path...)
			results = append(results,copied)
		}

		path = path[:l]
	}
	dfs(0,0)
	return results
}
