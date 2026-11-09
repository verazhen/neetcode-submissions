func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	path := []int{}
	results := [][]int{}
	var dfs func(i int, sum int)
	dfs = func(i int, sum int){
		if sum == target{
			copied := append([]int{},path...)
			results = append(results,copied)
			return
		}else if i == len(candidates){
			return
		}

		for j:=i; j <len(candidates);j++{

			if j>i && candidates[j] == candidates[j-1]{
				continue
			}else if sum+candidates[j] > target {
				break
			}
			path = append(path, candidates[j])

			dfs(j+1,sum+candidates[j])
			path = path[:len(path)-1]
		}
	}
	dfs(0,0)
	return results
}
