func findDuplicate(nums []int) int {
    // if traversing the array like link list
	// next next nodex to travers is the index with value of previous element
	// if more than one element point to the same node
	// that means we'll have a cycle link list
	// (0,1) -> (1,2) -> V(2,3) -> (3,2)
	// 0,0
	// 1,2
	// 2,2
	// (0,1) -> (1,2) -> (2,3) -> (3,4) -> V(4,4)
	// 0,0
	// 1,2
	// 2,4
	// 3,4
	// 4,4
	// index of the start node in cycle will be the answer
	// (0,1) -> (1,3) -> (3,2) -> V(2,4) -> (4,2)
	// 0,0
	// 1,3
	// 3,4
	// 2,4
	// 4,4


	slow := 0
	fast := 0

	for (slow != fast) || slow == 0{
		slow = nums[slow]
		fast = nums[nums[fast]]
	}

	slow = 0
	for (slow != fast) || slow == 0{
		slow = nums[slow]
		fast = nums[fast]
	}

	return slow
}
