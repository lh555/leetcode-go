package main

func maxProfit(price []int) int {
	if price == nil {
		return 0
	}
	minPrice := 0                   // 当前遇到的最低价格的"下标"
	dp := make([]int, len(price)+1) // dp[i] = 考虑前 i 天的最大利润
	dp[0] = 0                       // 0 天，无法交易
	dp[1] = 0                       // 1 天，无法买卖（买入和卖出不能是同一天）
	for i := 2; i <= len(price); i++ {
		if price[minPrice] > price[i-1] {
			minPrice = i - 1 // 更新最低价下标
		}
		dp[i] = max(dp[i-1], price[i-1]-price[minPrice])
		/*
			dp[i] 的两种选择：
			不卖：dp[i-1]，沿用之前的最大利润
			今天卖：price[i-1] - price[minPrice]，用当前价减去历史最低价
		*/
	}
	return dp[len(price)]
}
