package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	if nums == nil || len(nums) == 0 {
		return nil
	}
	a_map := make(map[int]int, len(nums))
	res := make([]int, 2)
	for i := 0; i < len(nums); i++ {
		if index, ok := a_map[target-nums[i]]; ok {
			res[0] = index
			res[1] = i
			//sort.Ints(res)
			return res
		}
		a_map[nums[i]] = i
		fmt.Println("i = ", i, "target-nums[i] = ", target-nums[i], "a_map = ", a_map)
	}
	return nil
}

func main() {
	res := twoSum([]int{2, 7, 11, 15}, 17)
	fmt.Println(res)
}

/*
sort.Ints 是一个原地排序函数，它没有返回值（返回 void），因此不能直接作为函数的返回值。
*/
