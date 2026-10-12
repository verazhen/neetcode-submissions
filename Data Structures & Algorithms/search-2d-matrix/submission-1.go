func searchMatrix(matrix [][]int, target int) bool {

	i := 0
	j := len(matrix) -1
	row := -1
	
	for i <= j {
		k := (i+j)/2
		l := len(matrix[k])
		min := matrix[k][0]
		max := matrix[k][l-1]

		if target >= min && target <=max{
			row = k
			break
		}else if target >= max{
			i = k+1
		}else{
			j = k-1
		}
	}

	if row == -1{
		return false
	}

	i = 0
	j = len(matrix[row]) -1
	for i <= j {
		k := (i+j)/2

		if target == matrix[row][k]{
			return true
		}else if target > matrix[row][k]{
			i = k+1
		}else{
			j = k-1
		}
	}
	return false
}
