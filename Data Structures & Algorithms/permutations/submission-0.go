func permute(nums []int) [][]int {

	numMap := map[int]bool{}
	path := []int{}
	results := [][]int{}
	var dfs func()
	dfs = func(){
		if len(path) == len(nums){
			copied := append([]int{},path...)
			results = append(results,copied)
			return
		}
		for _,num:=range nums{
			if ok:=numMap[num];ok{
				continue
			}
			numMap[num] = true
			path = append(path,num)
			dfs()
			numMap[num] = false
			path = path[:len(path)-1]
		}

	}
	dfs()
	return results
}
