func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	n := len(nums1)
	m := len(nums2)
	leftSize := (n + m + 1) / 2

	short := []int{}
	long := []int{}

	if n <= m {
		short = nums1
		long = nums2
	} else {
		short = nums2
		long = nums1
	}

	negInf := -1_000_001
	posInf := 1_000_001

	l := 0
	r := len(short)

	for l <= r {
		shortCut := (l + r) / 2
		longCut := leftSize - shortCut

		shortLeft := 0
		if shortCut == 0 {
			shortLeft = negInf
		} else {
			shortLeft = short[shortCut-1]
		}

		shortRight := 0
		if shortCut == len(short) {
			shortRight = posInf
		} else {
			shortRight = short[shortCut]
		}

		longLeft := 0
		if longCut == 0 {
			longLeft = negInf
		} else {
			longLeft = long[longCut-1]
		}

		longRight := 0
		if longCut == len(long) {
			longRight = posInf
		} else {
			longRight = long[longCut]
		}

		if shortLeft > longRight {
			// short 左半邊拿太多，切口往左
			r = shortCut - 1
		} else if longLeft > shortRight {
			// short 左半邊拿太少，切口往右
			l = shortCut + 1
		} else {
			// 找到合法 partition
			if (n+m)%2 == 0 {
				leftMax := max(shortLeft, longLeft)
				rightMin := min(shortRight, longRight)
				return (float64(leftMax) + float64(rightMin)) / 2
			}

			return float64(max(shortLeft, longLeft))
		}
	}

	return -1
}