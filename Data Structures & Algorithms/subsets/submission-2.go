//Recursive DFS + push/pop
func subsets(nums []int) [][]int {
    result := [][]int{}
    path := []int{}

    var dfs func(index int)
    dfs = func(index int) {
		if index >= len(nums){
			copied := append([]int{}, path...)
			result = append(result,copied)
			return
		}
    	path = append(path,nums[index])
		dfs(index+1)
		path = path[0:len(path)-1]
		dfs(index+1)
    }

    dfs(0)
    return result
}