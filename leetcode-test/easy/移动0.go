package main

import "fmt"

// 把 0 视作空位。我们要把所有非零元素都移到数组左边的空位上，并保证非零元素的顺序不变。
func moveZeroes(nums []int) {
	i0 := 0 // i0左边的元素都是最终位置
	for i, x := range nums {
		fmt.Println("交换前：", "i0 = ", i0, "i = ", i, "x = ", x, "nums = ", nums)
		if x != 0 {
			nums[i], nums[i0] = nums[i0], nums[i]
			i0++
			fmt.Println("交换后：", "i0 = ", i0, "i = ", i, "x = ", x, "nums = ", nums)
			continue
		}
		fmt.Println("不交换")
	}
}

func main() {
	nums := []int{2, 5, 0, 1, 0, 3, 12}
	moveZeroes(nums)
}
