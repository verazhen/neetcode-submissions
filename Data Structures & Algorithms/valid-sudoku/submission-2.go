func isValidSudoku(board [][]byte) bool {
	lineExistArr   := [9]int{}
	colExistArr    := [9]int{}
	squareExistArr := [9]int{}

	for i,line:=range board{
		for j,b:=range line{
			if b == '.'{
				continue
			}
			k:=getSquareIndex(i,j)

			mask:=1 << (b-'1')


			if lineExistArr[i] & mask != 0{
				return false
			}else if colExistArr[j] & mask != 0{
				return false
			}else if squareExistArr[k] & mask != 0{
				return false
			}

			lineExistArr[i] = lineExistArr[i] | mask
			colExistArr[j] = colExistArr[j] | mask
			squareExistArr[k] = squareExistArr[k] | mask
		}
	}

	return true
}

func getSquareIndex (i,j int)int{
	return i/3*3 + j/3
}

