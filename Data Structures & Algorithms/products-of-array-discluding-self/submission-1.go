func productExceptSelf(nums []int) []int {
	n := len(nums)
	result := make([]int, n)

	// Store the product of all elements to the left of i.
	leftProduct := 1
	for i := 0; i < n; i++ {
		result[i] = leftProduct
		leftProduct *= nums[i]
	}

	// Multiply by the product of all elements to the right of i.
	rightProduct := 1
	for i := n - 1; i >= 0; i-- {
		result[i] *= rightProduct
		rightProduct *= nums[i]
	}

	return result
}