package main

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	dp := make([]int, n+1)
	// 从1开始，如果从0开始，那dp[0] = 1，但是不好理解0阶楼梯的方法数。但是从1开始，需要处理n=1的情况，避免dp[2]超范围
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n; i++ {
		// 到达第i级 = 从i-1级爬1步 + 从i-2级一次爬2步
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}
