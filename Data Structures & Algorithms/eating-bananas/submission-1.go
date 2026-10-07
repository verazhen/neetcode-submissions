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
			//往前推 k-1 格，讓任何「有餘數」的除法都進入下一個整數區間，但又不會讓「剛好整除」的數字多進一個區間。
			total += (pile + k - 1) / k
		}
		if total > h {
			i = k + 1
		}else{
			j = k 
		}

	}
	return -1
}
