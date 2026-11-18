func generateParenthesis(n int) []string {
	str := ""
	results := []string{}
	leftCount := 0
	var dfs func()
	dfs = func(){
		if len(str) == n * 2{
			results = append(results,str)
			return
		}
		if leftCount < n{
			str = fmt.Sprintf("%s(",str)
			leftCount++
			dfs()
			str = str[:len(str)-1]
			leftCount--
		}

		if len(str)-leftCount < leftCount {
			str = fmt.Sprintf("%s)",str)
			dfs()
			str = str[:len(str)-1]
		}
	}
	dfs()
	return results
}
