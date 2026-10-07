func isValidSudoku(board [][]byte) bool {
	lineExistArr:= [9]map[int]struct{}{}
	colExistArr:= [9]map[int]struct{}{}
	squareExistArr:= [9]map[int]struct{}{}

	for i:=0;i <9;i++{
			lineExistArr[i] = map[int]struct{}{}
			colExistArr[i] = map[int]struct{}{}
			squareExistArr[i] = map[int]struct{}{}
	}


	for i,line:=range board{
		for j,content:=range line{
			if string(content) == "."{
				continue
			}

			num,err:=strconv.Atoi(string(content))
			if err!=nil{
				continue
			}
			k:=getSquareIndex(i,j)

			if _,ok:=lineExistArr[i][num];ok{
				return false
			}else if _,ok:=colExistArr[j][num];ok{
				return false
			}else if _,ok:=squareExistArr[k][num];ok{
				return false
			}

			lineExistArr[i][num] = struct{}{}
			colExistArr[j][num] = struct{}{}
			squareExistArr[k][num] = struct{}{}
		}
	}

	return true
}

func getSquareIndex (i,j int)int{
	return i/3*3 + j/3
}

