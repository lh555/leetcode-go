package main

import "fmt"

func twoSum2(nums []int, target int) []int {
	prevNums := map[int]int{}  // 创建了一个空的 map，key 是 int，value 也是 int
	for i, num := range nums {
		targetNum := target - num
		targetNumIndex, ok := prevNums[targetNum]
		if ok {
			return []int{targetNumIndex, i}
		} else {
			prevNums[num] = i
		}
	}
	return []int{}
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	result := twoSum2(nums, target)
	fmt.Println(result)
}
