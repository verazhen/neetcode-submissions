func evalRPN(tokens []string) int {
	//if operater : do calculation with top 2 ele in stack
	//if operand: push to stack
	resultStack:=[]int{}
	for _,token:=range tokens{
		if num, err := strconv.Atoi(token);err==nil{
			resultStack = append(resultStack,num)
			continue
		}

		last := len(resultStack)-1
		top := resultStack[last]
		resultStack = resultStack[:last]
		switch token {
			case "+":
				resultStack[last-1] += top
			case "-":
				resultStack[last-1] -= top
			case "*":
				resultStack[last-1] *= top
			case "/":
				resultStack[last-1] /= top
			default:
				
		}
	}

	if len(resultStack) == 0{
		return 0
	}

	return resultStack[0]
}
