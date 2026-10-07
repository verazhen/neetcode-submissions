func minEatingSpeed(piles []int, h int) int {
	//possible hour option: 1~ max pile
	//binary search hour option to check if its availale

	//1. find largest pile by iterating the piles
	maxPile := 0
	for _, pile:=range piles{
		maxPile = max(maxPile, pile)
	}

	//binary search
	i:= 1
	j:= maxPile
	for i <= j{
		if i == j{
			return i
		}

		k := (i+j) /2
		//iterate piles to calcuate hours
		total := 0
		for _, pile :=range piles{
			total += pile / k
			if pile % k > 0{
				total++
			}
		}
		if total > h {
			i = k + 1
		}else{
			j = k 
		}

	}
	return -1
}
