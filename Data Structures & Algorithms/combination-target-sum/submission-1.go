func combinationSum(nums []int, target int) [][]int {
    path:=[]int{}
	results:=[][]int{}
	var dfs func(i int,sum int)
	dfs = func(i int,sum int){
		if sum == target{
			//copy result to path
			copied := append([]int{},path...)
			results = append(results,copied)
			return
		}else if i >= len(nums)|| sum> target{
			return
		}
		
		//include
		//選擇該數，接下來可自由選擇要不要再選一次
		path = append(path,nums[i])
		dfs(i,sum+nums[i])

		//exclude
		//此數值沒選的話，跳到下一個數字
		path = path[:len(path)-1]
		dfs(i+1,sum)
	}
	dfs(0,0)
	return results
}
