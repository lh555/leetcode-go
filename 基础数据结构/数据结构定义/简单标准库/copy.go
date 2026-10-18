package main

import "fmt"

func copySlice(a []int, i int, x int, n int) []int {
	// 删除a[i]，可以用 copy 将i+1到末尾的值覆盖到i,然后末尾-1
	copy(a[i:], a[i+1:])
	a = a[:len(a)-1]
	fmt.Println("删除a[i]后:", a)
	// make创建长度，则通过索引赋值
	a = make([]int, n)
	a[n-1] = x
	a = append(a, x)
	fmt.Println("make创建长度，则通过索引赋值:", a)
	// make长度为0，则通过append()赋值
	a = make([]int, 0)
	fmt.Println("make长度为0，则通过append()赋值前:", a)
	a = append(a, x)
	fmt.Println("make长度为0，则通过append()赋值:", a)
	return a
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	// 删除a[1]，将3赋值给a[1]
	copySlice(a, 1, 3, 6)
}
