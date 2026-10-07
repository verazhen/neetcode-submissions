func checkInclusion(s1 string, s2 string) bool {
	expectWord := map[byte]int{}
	expectLen := 0
	for i:=range s1{
		expectLen++
		expectWord[s1[i]] ++
	}

	i,j:=0,0
	for j < len(s2){
		s := s2[j]
		expectWord[s] --

		for i <= j{
			if count:=expectWord[s];count>=0 {
				break
			}
			expectWord[s2[i]] ++
			i++
		}

		
		if j-i+1 == expectLen{
			return true
		}
		j++
	}
	return false
}
