func lengthOfLongestSubstring(s string) int {
	wordMap:=map[byte]bool{}
	i,j:=0,0
	maxLen:=0
	for j<len(s){
		add:=s[j]
		for i<j {
			if ok :=wordMap[add];!ok{
				break
			}
			removed:=s[i]
			wordMap[removed] = false
			i++
		}
		wordMap[add] = true
		maxLen = max(maxLen,j-i+1)
		j++
	} 
	return maxLen
}

