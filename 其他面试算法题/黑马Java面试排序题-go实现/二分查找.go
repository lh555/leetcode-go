package main

import "fmt"

func binarySearch(a []int, t int) int {
	l, r := 0, len(a)-1
	var m int
	for l <= r {
		m = (l + r) / 2 // l+r 可能会溢出，使用 l+(r-l)/2 避免溢出，也可以用 l+(r-l)>>1 位运算避免溢出，正整数才能用
		if a[m] == t {
			return m
		} else if a[m] > t {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return -1
}

// 递归
func recursion(a []int, l, r, t int) int {
	if l > r {
		return -1
	}
	m := l + (r-l)/2
	if a[m] == t {
		return m
	}
	if a[m] > t {
		return recursion(a, l, m-1, t)
	} else {
		return recursion(a, m+1, r, t)
	}
}

// 递归写法2
//func recursion(a []int, l, r, t int) int {
//	if l > r {
//		return -1
//	}
//	m := l + (r-l)/2
//	if a[m] == t {
//		return m
//	}
//	var p1 int
//	var p2 int
//	if a[m] > t {
//		p1 = recursion(a, l, m-1, t)
//	} else {
//		p2 = recursion(a, m+1, r, t)
//	}
//	if p1 != -1 && p1 != 0 {
//		return p1
//	} else {
//		return p2
//	}
//}

func main() {
	a := []int{1, 5, 8, 11, 19, 22, 31, 35, 40, 45, 48, 49, 50}
	t := 2
	res := binarySearch(a, t)
	fmt.Println(res)

	// 递归
	res2 := recursion(a, 0, len(a)-1, 5)
	fmt.Println(res2)
}
