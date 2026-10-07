func evalRPN(tokens []string) int {
	//if operater : do calculation with top 2 ele in stack
	//if operand: push to stack
	resultStack:=[]int{}
	for _,token:=range tokens{
		last := len(resultStack)-1
		switch token {
			case "+":
				//pop and cal
				top := resultStack[last]
				resultStack = resultStack[:last]
				resultStack[last-1] += top
			case "-":
				top := resultStack[last]
				resultStack = resultStack[:last]
				resultStack[last-1] -= top
			case "*":
				top := resultStack[last]
				resultStack = resultStack[:last]
				resultStack[last-1] *= top
			case "/":
				top := resultStack[last]
				resultStack = resultStack[:last]
				resultStack[last-1] /= top
			default:
				num, _ := strconv.Atoi(token)
				resultStack = append(resultStack,num)
		}
	}

	if len(resultStack) == 0{
		return 0
	}

	return resultStack[0]
}
