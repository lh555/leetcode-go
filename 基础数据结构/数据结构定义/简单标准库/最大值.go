package main

import (
	"fmt"
	"math"
)

func main() {
	// int32 最大最小值
	max32 := math.MaxInt32 // 实际值：1<<31-1
	min32 := math.MinInt32 // 实际值：-1<<31
	// int64 最大最小值（int默认是int64）
	max64 := math.MaxInt64 // 实际值：1<<63-1
	min64 := math.MinInt64 // 实际值：-1<<63
	// 打印结果
	fmt.Println("int32 最大最小值:", max32, min32)
	fmt.Println("int64 最大最小值:", max64, min64)
}
