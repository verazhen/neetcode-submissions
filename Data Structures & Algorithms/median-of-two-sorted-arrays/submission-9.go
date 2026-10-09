func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	n := len(nums1)
	m := len(nums2)
	leftSize := (n+m+1)/2
	numA := []int{} //shorter
	numB := []int{}
	if n <=m {
		numA = nums1
		numB = nums2
	}else{
		numA = nums2
		numB = nums1
	}
	negMax := -1000001
	posMax := 1000001

	l:= 0
	r:= len(numA)
	for l <= r{
		numALeft := (l+r+1)/2 
		numBLeft := leftSize - numALeft
		
		aLeft:= 0
		if numALeft == 0 {
			aLeft = negMax
		}else{
			aLeft = numA[numALeft - 1]
		}
		aRight:= 0
		if numALeft == len(numA) {
			aRight = posMax
		}else{
			aRight = numA[numALeft]
		}

		bLeft:= 0
		if numBLeft == 0 {
			bLeft = negMax
		}else{
			bLeft = numB[numBLeft - 1]
		}

		bRight:= 0
		if numBLeft == len(numB) {
			bRight = posMax
		}else{
			bRight = numB[numBLeft]
		}

		fmt.Println(numALeft,aLeft,aRight,bLeft,bRight)
		if aLeft > bRight{
			//move left
			r = numALeft  - 1
		}else if bLeft > aRight{
			//move right
			l = numALeft  + 1
		}else{
			//calculate result
			if (n+m)%2 == 0 {
				return (float64(max(aLeft,bLeft))+float64(min(aRight,bRight)))/2
			}
			return float64(max(aLeft,bLeft))
		}
		
	}
	return -1
}



