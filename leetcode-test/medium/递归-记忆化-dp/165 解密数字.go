package main

import (
	"fmt"
	"strconv"
)

func crackNumber(num int) int {
	str := strconv.Itoa(num) // 将整数转为字符串
	return process([]rune(str), 0)
}

// process 处理从索引 i 开始的字符串，返回可能的翻译方法数
func process(c []rune, i int) int {
	// 如果已经处理完所有字符，说明找到一种有效翻译
	if i == len(c) {
		return 1
	}

	// 如果当前字符是 '0'，无法单独翻译，返回0
	if c[i] == '0' {
		return 0
	}

	// 可能性1：当前字符单独翻译
	ways := process(c, i+1)

	// 可能性2：当前字符和下一个字符组合翻译（两位数字）
	if i+1 < len(c) {
		// 计算两位数
		twoDigit := int(c[i]-'0')*10 + int(c[i+1]-'0')
		if twoDigit < 26 {
			ways += process(c, i+2)
		}
	}

	return ways
}

// 记忆化搜索
func crackNumber2(num int) int {
	str := strconv.Itoa(num)
	// 创建长度为 len(str)+1 的缓存，因为 i 可以取到 len(str)
	cache := make([]int, len(str)+1)
	return process2([]rune(str), 0, cache)
}

func process2(c []rune, i int, cache []int) int {
	// 先判断是否到达末尾，避免访问 cache[len(c)] 越界
	if i == len(c) {
		return 1
	}
	if c[i] == '0' {
		return 0
	}
	// 检查缓存
	if cache[i] != 0 {
		return cache[i]
	}
	ways := process(c, i+1)
	if i+1 < len(c) {
		twoDigit := int(c[i]-'0')*10 + int(c[i+1]-'0')
		if twoDigit < 26 {
			ways += process(c, i+2)
		}
	}
	cache[i] = ways
	return ways
}

// dp
func crackNumber3(num int) int {
	str := strconv.Itoa(num)
	return process3(str)
}

func process3(str string) int {
    n := len(str)
    if n == 0 {
        return 0
    }
    
    // dp[i] 表示从 i 到末尾的翻译方法数
    // 对应递归中的 process(c, i)
    dp := make([]int, n+1)
    
    // 递归终止条件：i == len(c) 返回 1
    dp[n] = 1
    
    // 从后往前计算，对应递归的递推过程
    for i := n - 1; i >= 0; i-- {
        // 对应递归中的 if c[i] == '0' { return 0 }
        if str[i] == '0' {
            dp[i] = 0
            continue
        }
        
        // 对应递归中的 ways := process(c, i+1)
        ways := dp[i+1]
        
        // 对应递归中的 if i+1 < len(c) && twoDigit < 26 { ways += process(c, i+2) }
        if i+1 < n {
            twoDigit := int(str[i]-'0')*10 + int(str[i+1]-'0')
            if twoDigit < 26 {
                ways += dp[i+2]
            }
        }
        
        dp[i] = ways
    }
    
    return dp[0]
}

func main() {
	res := crackNumber(216612)
	fmt.Println(res)

	res2 := crackNumber2(216612)
	fmt.Println(res2)
	res3 := crackNumber2(216612)
	fmt.Println(res3)
}

/*
字符串处理：使用 []rune 而不是 []char，因为 Go 中 rune 代表 Unicode 码点

显式类型转换：Go 需要显式类型转换，如 int(c[i]-'0')
*/
