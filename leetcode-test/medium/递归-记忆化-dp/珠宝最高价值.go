package main

// 暴力递归
func jewelleryValue(frame [][]int) int {
	if frame == nil || len(frame) == 0 || len(frame[0]) == 0 {
		return 0
	}
	return recur(frame, 0, 0)
}

func recur(frame [][]int, i, j int) int {
	if i >= len(frame) || j >= len(frame[0]) {
		return 0
	}
	down := recur(frame, i+1, j) + frame[i][j]
	right := recur(frame, i, j+1) + frame[i][j]
	return max(right, down)
}

// 记忆化
func jewelleryValue2(frame [][]int) int {
	if frame == nil || len(frame) == 0 || len(frame[0]) == 0 {
		return 0
	}
	cache := make([][]int, len(frame))
	for i := range cache {
		cache[i] = make([]int, len(frame[0]))
	}
	for i := 0; i < len(cache); i++ {
		for j := 0; j < len(cache[0]); j++ {
			cache[i][j] = -1
		}
	}
	return recur(frame, 0, 0, cache)
}

func recur2(frame [][]int, i, j int, cache [][]int) int {
	if i >= len(frame) || j >= len(frame[0]) {
		return 0
	}
	if cache[i][j] != -1 {
		return cache[i][j]
	}
	right := recur(frame, i, j+1, cache) + frame[i][j]
	down := recur(frame, i+1, j, cache) + frame[i][j]
	cache[i][j] = max(right, down)
	return cache[i][j]
}

// 严格表结构 dp
func jewelleryValue3(frame [][]int) int {
	if frame == nil || len(frame) == 0 || len(frame[0]) == 0 {
		return 0
	}
	dp := make([][]int, len(frame)+1) // 一般比原数组多一个长度即可
	for i := range dp {
		dp[i] = make([]int, len(frame[0])+1)
	}
	// 初始化dp数组
	for i := 0; i < len(dp); i++ {
		dp[i][len(dp[0])-1] = 0 // 这一步其实可以不要，初始化dp切片的时候，就是0了
	}
	for j := 0; j < len(dp[0]); j++ {
		dp[len(dp)-1][j] = 0 // 这一步其实可以不要，初始化dp切片的时候，就是0了
	}
	return recur3(frame, 0, 0, dp)
}

func recur3(frame [][]int, i, j int, dp [][]int) int {
	// 当前格子是由右边和下边的最大值得到的，所以是从右下角向左上角方向遍历，一个一个格子推出来
	// len(dp[0]) - 1 = len(frame)，会导致frame[i][j]越界，所以直接-2
	for i := len(dp) - 2; i >= 0; i-- {
		for j := len(dp[0]) - 2; j >= 0; j-- {
			// right := recur(frame, i, j+1, cache) + frame[i][j]
			// down := recur(frame, i+1, j, cache) + frame[i][j]
			// 其实就是把递归函数名称改为dp，如果去掉frame，cache这些切片，只保留下标i，j，这样就不用再画图了
			right := dp[i][j+1] + frame[i][j]
			down := dp[i+1][j] + frame[i][j]
			dp[i][j] = max(right, down)
		}
	}
	// 递归是往右、往下走，动态规划就是反过来，当前格子是由右边和下边的最大值得到的，也就是说格子越靠近左上角，越大
	return dp[0][0]
}
