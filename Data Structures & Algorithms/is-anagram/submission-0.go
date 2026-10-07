func isAnagram(s string, t string) bool {
	if len(s)!=len(t){
		return false
	}
	
	charList := [26]int{}
	for _,char:=range s{
		i:=char-'a'//rune index
		charList[i] ++
	}
		for _,char:=range t{
		i:=char-'a'//rune index
		charList[i] --
		if charList[i] <0{
			return false
		}
	}
	return true
}
