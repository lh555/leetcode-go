package main

import "fmt"

func main() {
	//创建队列
	queue := make([]int, 0)
	//enqueue入队
	queue = append(queue, 10)
	//dequeue出队
	v := queue[0]
	queue = queue[1:]
	fmt.Println("v = ", v)
	//检查队列为空
	if len(queue) == 0 {
		fmt.Println("队列为空")
	} else {
		fmt.Println(queue)
	}
}

/*
潜在问题
虽然时间复杂度是 O(1)，但这种实现方式有一个潜在的内存问题：

随着不断出队，原切片的底层数组可能会保留很多未使用的空间
因为切片的容量不会自动减少，即使队列变得很小，底层数组可能仍然很大
这可能导致内存使用效率不高，特别是在频繁操作的场景下
对比环形队列
相比之下，环形队列（ring buffer）的实现：
可以更有效地利用内存，因为它循环使用底层数组的空间
避免了切片操作可能带来的内存浪费
出队操作同样可以做到 O(1) 的时间复杂度
*/
