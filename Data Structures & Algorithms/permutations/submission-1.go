//不用hashmap記錄，改用swap把已選擇的放array前半段
func permute(nums []int) [][]int {
	results := [][]int{}
	var dfs func(i int)
	//i 是還沒被選擇的後半段開頭
	dfs = func(i int){
		if i == len(nums){
			copied := append([]int{},nums...)
			results = append(results,copied)
			return
		}
		for j:=i; j< len(nums);j++{
			nums[i],nums[j] = nums[j],nums[i]
			dfs(i+1)
			nums[i],nums[j] = nums[j],nums[i]
		}

	}
	dfs(0)
	return results
}
