package main

import "fmt"

func ways1(start int, K int, aim int, N int) int {
	if N < 2 || start < 1 || start > N || aim < 1 || aim > N || K < 1 {
		return -1
	}
	return process1(start, K, aim, N)
}

// 暴力递归   机器人在一个一维数组上走K步
// 机器人当前来到的位置是cur，
// 机器人还有rest步需要去走，
// 最终的目标是aim，
// 有哪些位置？1~N
// 返回：机器人从cur出发，走过rest步之后，最终停在aim的方法数，是多少？
func process1(cur int, rest int, aim int, N int) int {
	if rest == 0 { // 如果已经不需要走了，走完了！
		if cur == aim {
			return 1 //走完后如果当前位置是终点，那么该方法是对的
		}
		return 0 
	}
	// (cur, rest)
	if cur == 1 { // 1 -> 2  当前在1位置，只能是往2走，没其他选择
		return process1(2, rest-1, aim, N)
	}
	// (cur, rest)
	if cur == N { // N-1 <- N   当前在N位置，只能是往N-1走，没其他选择
		return process1(N-1, rest-1, aim, N)
	}
	// (cur, rest)    机器人来到中间位置，此时会有两种选择
	return process1(cur-1, rest-1, aim, N) + process1(cur+1, rest-1, aim, N)
}

// 第一个优化：记忆化搜索，就是在暴力递归上加个缓存
func ways2(start int, K int, aim int, N int) int {
	if N < 2 || start < 1 || start > N || aim < 1 || aim > N || K < 1 {
		return -1
	}
	// 有2个可变参数，故用二维表,K范围是0~K，aim是1~N,求范围只是省空间，二维表可以无限大，空间够用即可
	cache := make([][]int, N+1)
	for i := 0; i <= N; i++ {
		cache[i] = make([]int, K+1)
		for j := 0; j <= K; j++ {
			cache[i][j] = -1 // -1表示这个状态没有计算过
		}
	}
	// cache就是缓存表
	// cache[cur][rest] == -1 -> process1(cur, rest)之前没算过！
	// cache[cur][rest] != -1 -> process1(cur, rest)之前算过！返回值，cache[cur][rest]
	// N+1 * K+1
	return process2(start, K, aim, N, cache)
}

// cur 范: 1 ~ N
// rest 范：0 ~ K
func process2(cur int, rest int, aim int, N int, cache [][]int) int {
	if cache[cur][rest] != -1 { // 如果这个状态已经算过了，直接返回这个格子的值
		return cache[cur][rest]
	}
	// 之前没算过！
	if rest == 0 { // 如果已经不需要走了，走完了！
		if cur == aim {
			cache[cur][rest] = 1 // 缓存没有命中，返回前先记录这个格子的值表明已经计算过这个格子了
			return 1 //走完后如果当前位置是终点，那么该方法是对的
		}
		cache[cur][rest] = 0
		return 0
	}
	// (cur, rest)
	if cur == 1 { // 1 -> 2  当前在1位置，只能是往2走，没其他选择
		ans := process2(2, rest-1, aim, N, cache)
		cache[cur][rest] = ans
		return ans
	}
	// (cur, rest)
	if cur == N { // N-1 <- N   当前在N位置，只能是往N-1走，没其他选择
		ans := process2(N-1, rest-1, aim, N, cache)
		cache[cur][rest] = ans
		return ans
	}
	// (cur, rest)    机器人来到中间位置，此时会有两种选择
	ans := process2(cur-1, rest-1, aim, N, cache) + process2(cur+1, rest-1, aim, N, cache)
	cache[cur][rest] = ans
	return ans
}

// 严格表结构dp
func ways3(start int, K int, aim int, N int) int {
	if N < 2 || start < 1 || start > N || aim < 1 || aim > N || K < 1 {
		return -1
	}
	dp := make([][]int, N+1)  // 以 dp 表来看，行是 cur（上小下大），列是 rest（左小右大）
	for i := 0; i <= N; i++ {
		dp[i] = make([]int, K+1)
	}
	dp[aim][0] = 1 // rest==0，并且在终点处，方法是1
	for rest := 1; rest <= K; rest++ { // 从后往前推，上边暴力递归和记忆化搜索是从前往后推
		dp[1][rest] = dp[2][rest-1] // 依赖左下角   当前位置是1时，只能右走，直接获取左下角值
		for cur := 2; cur < N; cur++ { // 一般情况
			// 依赖左上角和左下角
			dp[cur][rest] = dp[cur-1][rest-1] + dp[cur+1][rest-1]
		}
		dp[N][rest] = dp[N-1][rest-1] // 依赖左上角  当前位置是N时，只能左走，直接获取左上角值
	}
	return dp[start][K]
}

func main() {
	fmt.Println(ways1(2, 6, 4, 5))
	fmt.Println(ways2(2, 6, 4, 5))
	fmt.Println(ways3(2, 6, 4, 5))
}