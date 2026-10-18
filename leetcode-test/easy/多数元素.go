package main

import "fmt"

// 使用map做
func majorityElement(nums []int) int {
	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
	}
	max_k := 0     // 记录出现次数最多的元素值
	max_count := 0 // 记录最大出现次数
	for k, v := range m {
		if max_count < v {
			max_count = v
			max_k = k
		}
	}
	return max_k
}

func main() {
	nums := []int{1, 2, 2, 3, 3, 3, 4}
	res := majorityElement(nums)
	fmt.Println(res)
}
