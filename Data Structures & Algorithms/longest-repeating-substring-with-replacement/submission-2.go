func characterReplacement(s string, k int) int {
	i,j := 0,0
	freqMap:=map[byte]int{}
	maxLen:=0
	maxFreqInWindow:=0
	for j < len(s){
		//join right pointer
		freqMap[s[j]]++
		maxFreqInWindow = max(maxFreqInWindow,freqMap[s[j]])
		//check if valid
		for (j-i+1) - maxFreqInWindow > k{
			//if invalid, remove left til valid and 
			freqMap[s[i]] --
			//re-cal maxFreq in case maxFreq decrease
			if freqMap[s[i]] +1 == maxFreqInWindow{
					newMax := 0
				for _,count:=range freqMap{
					newMax = max(newMax,count)
				}
				maxFreqInWindow = newMax
			}
			i++
		}
		maxLen = max(maxLen,(j-i+1))
		j++
	}
	return maxLen
}
