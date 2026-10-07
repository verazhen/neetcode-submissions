func trap(height []int) int {
	n := len(height)

	// rightMaxArr[i] = max height from i to n-1
	rightMaxArr := make([]int, n)
	rightMax := 0

	for i := n - 1; i >= 0; i-- {
		rightMax = max(rightMax, height[i])
		rightMaxArr[i] = rightMax
	}

	leftMax := 0
	water := 0

	for i := 0; i < n; i++ {
		leftMax = max(leftMax, height[i])

		waterLevel := min(leftMax, rightMaxArr[i])
		water += waterLevel - height[i]
	}

	return water
}