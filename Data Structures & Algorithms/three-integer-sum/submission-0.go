func threeSum(nums []int) [][]int {
	//indices i, j and k are all distinct.
	//output should not contain any duplicate triplets
	//hashmap 在固定排序規則後，可以記錄 unique triplets，但是如果要記錄三個值是不同index需要另外處理
	//sort + two pointers，可以夾擊計算左右指針總和 也可以做到 distinct index 跟 triplet 的dedup
	//O(nlogn)，不確定golang預設是什麼排序法
	sort.Ints(nums)
	//nums[i]+nums[j] == nums[-k] * -1
	k:=0

	results := [][]int{}
	for k < len(nums){
		if k!=0 && nums[k] == nums[k-1]{
			k++
			continue
		}
		target := nums[k] * -1
		i:=k+1
		j:=len(nums)-1
		for i < j{

			if (i > k+1 && nums[i] == nums[i-1]){
				i++
				continue
			}else if j!=len(nums)-1 && nums[j] == nums[j+1]{
				j--
				continue
			}


			total := nums[i] + nums[j]
			if total == target {
				results = append(results,[]int{nums[i],nums[j],nums[k]})
				i++
				j--
			}else if total < target{
				i++
			}else{
				j--
			}
		}
		k++
	}
	return results

}
