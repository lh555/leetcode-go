package main

import (
	"fmt"
	"sort"
)

func sortSlice(sliceInt []int, sliceString []string) ([]int, []string) {
	// int排序
	sort.Ints(sliceInt)
	fmt.Println("int排序结果:", sliceInt)
	// 字符串排序
	sort.Strings(sliceString)
	fmt.Println("字符串排序结果:", sliceString)
	// 自定义排序
	sort.Slice(sliceString, func(i, j int) bool { return sliceString[i] > sliceString[j] })
	fmt.Println("自定义排序结果:", sliceString)
	return sliceInt, sliceString
}

func main() {
	sliceInt := []int{3, 2, 1}
	sliceString := []string{"c", "b", "a", "d"}
	sortSlice(sliceInt, sliceString)
}
