func characterReplacement(s string, k int) int {
	i,j := 0,0
	freqMap:=map[byte]int{}
	maxLen:=0
	maxFreq:=0
	for j < len(s){
		//join right pointer
		freqMap[s[j]]++
		maxFreq = max(maxFreq,freqMap[s[j]])
		//check if invalid, remove left and continue 
		if (j-i+1) - maxFreq > k{
			freqMap[s[i]] --
			i++
		}
		maxLen = max(maxLen,(j-i+1))
		j++
	}
	return maxLen
}
