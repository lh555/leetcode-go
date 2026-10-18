package main

import "fmt"

func main() {
	//创建栈
	stack := []int{1, 5, 3, 0}
	//push压入栈
	stack = append(stack, 10)
	fmt.Println(stack)
	//pop弹出
	v := stack[len(stack)-1]
	stack = stack[:len(stack)-1]
	fmt.Println("v = ", v)
	//检查栈空
	if len(stack) == 0 {
		fmt.Println("栈空")
	} else {
		fmt.Println(stack)
	}

}

// go没有队列，只能用slice模拟
