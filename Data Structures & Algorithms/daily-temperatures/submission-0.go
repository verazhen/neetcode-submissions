func dailyTemperatures(temperatures []int) []int {
	stack := [][2]int{} //value, index
	results := make([]int,len(temperatures))
	for j,t:=range temperatures{
		//if t > stack[last], calculate stack
		for len(stack)>0 {
			last := len(stack)-1
			top := stack[last][0]
			i := stack[last][1]
			if t <= top{
				break
			}
			stack = stack[:last]
			results[i] = j-i
		}
		//if t <= stack[last], push to stack only
		stack = append(stack,[2]int{t,j})
	}

	return results
}
