//include & exclude dfs
func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)

	path:=[]int{}
	results:=[][]int{}
	var dfs func(i int)
	dfs = func(i int){
		//copy result to path
		copied := append([]int{},path...)
		results = append(results,copied)
		if i >= len(nums){
			return
		}
		
		//[1,1,1,1,2,3]
		//要選擇哪些數字：接下來所有unique 的數字（包含i）
		value:=-21
		for j:=i;j<len(nums);j++{
			if nums[j] == value{
				continue
			}
			value = nums[j]
			path = append(path,nums[j])
			dfs(j+1)
			path = path[:len(path)-1]
		}
	}
	dfs(0)
	return results
}
