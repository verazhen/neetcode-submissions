func maxProfit(prices []int) int {
	//for every sell point, we only cares about the lowest buy point before
	lowestBuy:=prices[0]
	maxProfit:=0
	for i:=1;i<len(prices);i++{
		maxProfit = max(maxProfit,prices[i]-lowestBuy)
		lowestBuy = min(lowestBuy,prices[i])
	}
	return maxProfit
}
