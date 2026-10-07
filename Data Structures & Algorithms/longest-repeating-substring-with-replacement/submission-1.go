func characterReplacement(s string, k int) int {
	i,j := 0,0
	freqMap:=map[byte]int{}
	maxFreq:=0
	freq:=0
	for j < len(s){
		//join right pointer
		freqMap[s[j]]++
		freq = max(freq,freqMap[s[j]])
		//check if valid
		for i+k <= j{
			if (j-i+1) - freq <= k{
				break
			}
			//if invalid, remove left til valid and 
			freqMap[s[i]] --
			//re-cal maxFreq in case maxFreq decrease
			if freqMap[s[i]] +1 == freq{
					newMax := 0
				for _,count:=range freqMap{
					newMax = max(newMax,count)
				}
				freq = newMax
			}
			i++
		}
		maxFreq = max(freq,maxFreq)
		j++
	}
	return min(maxFreq+k,len(s))
}
