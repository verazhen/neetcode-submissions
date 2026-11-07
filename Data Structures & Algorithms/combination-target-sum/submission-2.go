//分支不是選或不選，而是可以選哪些數字
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
		
		for j:=i;j< len(nums);j++{
			path = append(path,nums[j])
			dfs(j,sum+nums[j])
			path = path[:len(path)-1]
		}
	}
	dfs(0,0)
	return results
}

