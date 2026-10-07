func lengthOfLongestSubstring(s string) int {
	wordMap:=map[byte]int{} //string to index
	i,j:=0,0
	maxLen:=0
	for j<len(s){
		add:=s[j]
		if latestIndex,ok :=wordMap[add];ok && latestIndex>=i{
			i = latestIndex+1
		}
		wordMap[add] = j
		maxLen = max(maxLen,j-i+1)
		j++
	} 
	return maxLen
}

