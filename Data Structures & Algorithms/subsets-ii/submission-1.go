//include & exclude dfs
func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)

	path:=[]int{}
	results:=[][]int{}
	var dfs func(i int)
	dfs = func(i int){
		if i >= len(nums){
			//copy result to path
			copied := append([]int{},path...)
			results = append(results,copied)
			return
		}
		
		//include
		//選擇第一個不重複值，接下來可自由選擇要不要選下一個數值
		path = append(path,nums[i])
		dfs(i+1)

		//exclude
		//此數值沒選的話，剩下相同數值都不用選直接跳到 j
		path = path[:len(path)-1]
		j := i + 1
		for j < len(nums) && nums[j] == nums[i] {
    		j++
		}
		dfs(j)
	}
	dfs(0)
	return results
}
