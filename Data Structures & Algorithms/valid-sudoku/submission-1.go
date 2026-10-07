func isValidSudoku(board [][]byte) bool {
	lineExistArr:= [9][9]bool{}
	colExistArr:= [9][9]bool{}
	squareExistArr:= [9][9]bool{}

	for i,line:=range board{
		for j,b:=range line{
			if b == '.'{
				continue
			}
			k:=getSquareIndex(i,j)

			if lineExistArr[i][b-'1']{
				return false
			}else if colExistArr[j][b-'1']{
				return false
			}else if squareExistArr[k][b-'1']{
				return false
			}

			lineExistArr[i][b-'1'] = true
			colExistArr[j][b-'1'] = true
			squareExistArr[k][b-'1'] = true
		}
	}

	return true
}

func getSquareIndex (i,j int)int{
	return i/3*3 + j/3
}

