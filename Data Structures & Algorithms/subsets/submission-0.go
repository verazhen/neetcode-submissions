func subsets(nums []int) [][]int {
	result := [][]int{}
	var dfs func (i int, path []int)
	dfs = func (i int, path []int){
		if i >= len(nums){
			copied := append([]int{}, path...)
			result = append(result,copied)
			return
		}
		
		dfs(i+1,append(path,nums[i]))
		dfs(i+1,path)
	}
	dfs(0,[]int{})
	return result
}
