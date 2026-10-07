func carFleet(target int, position []int, speed []int) int {
	carStack := [][2]float64{} //position to left time
	// calculate left time to target
	for i, pos:= range position{
		leftTime := float64(target-pos) / float64(speed[i])
		carStack = append(carStack,[2]float64{float64(pos),leftTime})
	}

	//sort cars from near to far from target (position decreasing)
	sort.Slice(carStack, func(i, j int) bool {
    	return carStack[i][0] > carStack[j][0]
	})

	// merge from l to r, if back left time > front, add to fleet stack
	lateFleet := carStack[0][1]
	fleet:=1
	for i:=1;i<len(carStack);i++{
		back := carStack[i][1]
		if back>lateFleet {
			//push back
			lateFleet = back
			fleet++
		}
	}

	return fleet
}
