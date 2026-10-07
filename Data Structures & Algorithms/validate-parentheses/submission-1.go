func isValid(s string) bool {
    stack:=[]rune{}
	parenthesesDict:=map[rune]rune{
		'(':')',
		'{':'}',
		'[':']',
	}
	for _,r:=range s{
		rb,ok:=parenthesesDict[r]
		//if left, push to stack
		if ok {
			stack = append(stack,rb)
			continue
		}
		//if right, check and pop from stack
		if len(stack) == 0 || stack[len(stack)-1] != r{
			return false
		}
		stack = stack[:len(stack)-1]
	}

	return len(stack)==0

}
