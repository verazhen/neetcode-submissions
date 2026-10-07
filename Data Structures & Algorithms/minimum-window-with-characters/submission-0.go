func minWindow(s string, t string) string {
	i,j := 0,0
	tMap:=map[byte]int{}
	for i:=range t{
		tMap[t[i]]++
	}
	remain:=len(t)
	minStart:=0
	minLen:=-1
	for j < len(s){
		if tMap[s[j]] > 0 {
			remain--
		}

			tMap[s[j]]--

		//check shortest valid window
		for remain == 0 && i<=j{ 
			if tMap[s[i]]==0{
				if minLen == -1 || j-i+1 < minLen{
					minStart = i
					minLen = j-i+1
				}
				break
			}
				tMap[s[i]]++
			i++ 
		}
		j++
	}
	if minLen < 0{
		return ""
	}

	return s[minStart:minStart+minLen]
}
