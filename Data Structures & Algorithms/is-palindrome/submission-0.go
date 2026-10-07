func isPalindrome(s string) bool {
	
	arr := []string{}
	for _,r := range s{
		if !(r >= 'a' && r <= 'z') && !(r >= 'A' && r <= 'Z') && !(r >= '0' && r <= '9'){
			continue
		}
		str := strings.ToLower(string(r))
		arr = append(arr,str)
	}
	halfLength := len(arr)/2 

	wordStack := []string{}
	for i,str := range arr{
		
		if i < halfLength{
			//push
			wordStack = append(wordStack, str)
		}else if len(arr)%2 == 1 && i == halfLength{
			//do nothing
		}else{
			//pop
			pop := wordStack[len(wordStack)-1]
			wordStack = wordStack[0:len(wordStack)-1]
			if str != pop {
				return false
			}
		}
		continue
	}
	return true
}
