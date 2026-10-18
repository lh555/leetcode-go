package main

import "fmt"

/*
**算法描述**
1. 依次比较数组中相邻两个元素大小，若 a[j] > a[j+1]，则交换两个元素，两两都比较一遍称为一轮冒泡，结果是让最大的元素排至最后
2. 重复以上步骤，直到整个数组有序
*/

func bubble(a []int) {
	// i < len(a)-1 是因为对于长度为 n 的数组，最多需要 n-1 轮排序，写成 len(a)也可以，但是 len(a)-1 更符合冒泡排序的逻辑
	for i := 0; i < len(a)-1; i++ {
		for j := 0; j < len(a)-1-i; j++ { // 优化点1：每经过一轮冒泡，内层循环就可以减少一次
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
}

/*
* 优化点1：每经过一轮冒泡，内层循环就可以减少一次
* 优化点2：如果某一轮冒泡没有发生交换，则表示所有数据有序，可以结束外层循环
 */

func bubbleOptimize(a []int) {
	for i := 0; i < len(a)-1; i++ {
		// 优化点2：如果某一轮冒泡没有发生交换，则表示所有数据有序，可以结束外层循环
		swap := false
		for j := 0; j < len(a)-1-i; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
				swap = true
			}
		}
		if !swap {
			break
		}
	}
}

func main() {
	a := []int{5, 1, 8, 9, 2, 8, 6, 7}
	bubble(a) // 底层数组共享，bubble 会改变 a 的值
	fmt.Println(a)
	b := []int{5, 1, 8, 9, 2, 8, 6, 7}
	bubbleOptimize(b)
	fmt.Println(b)
}
