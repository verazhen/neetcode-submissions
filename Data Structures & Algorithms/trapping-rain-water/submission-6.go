func trap(height []int) int {
	//由後往前找到當前的最高值
	//[3,3,3,3,3,3,3,3,2,1]
	//[3,2,1,5] -> [3,5,5,5]
	//[3,2,4,5] -> [3,5,5,5]
	maxRight:=0
	maxArr:=make([]int, len(height))
	for j:=len(height)-1;j>=0;j--{
		maxRight = max(maxRight,height[j])
		maxArr[j] = maxRight
	}


	//結算資訊，需同時符合以下條件才能記錄水位：
	//找到高於該bar 的左側bar高度
	//找到右側bar
	//1. 右側比左側高
	//2. 右側為右側窗口中最高
	i:=0
	total:=0
	for i < len(height)-2{
		if height[i] == 0 {
			i++
			continue
		}

		//確定 i 找k
		rightMax := maxArr[i+1]
		k:=i+1
		notFound:=false
		for k < len(height){
			if height[k] >= height[i]{
				break
			}else if height[k] == rightMax{
				break
			}
			//not graceful
			if k == len(height)-1{
				notFound = true
				break
			}
			k++
		}
		if notFound{
			i++
			continue
		}
		//結算j bar 水位
		for j:=i+1; j<k;j++{
			total += max(min(height[i],height[k]) - height[j],0)
		}

		i= k 
	}

	return total
}
