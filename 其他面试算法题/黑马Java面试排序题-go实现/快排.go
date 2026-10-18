package main

import "fmt"

func quick(a []int, l int, r int) {
	if l >= r {
		return
	}
	p := partition(a, l, r)
	quick(a, l, p-1)
	quick(a, p+1, r)
}

func partition(a []int, l int, r int) int {
	pv := a[l]
	i := l
	j := r
	for i < j {
		// j 从右找小的
		for i < j && a[j] > pv {
			j--
		}
		// i 从左找大的
		for i < j && a[i] <= pv {
			i++
		}
		a[i], a[j] = a[j], a[i]
	}
	a[l], a[j] = a[j], a[l]
	fmt.Printf("pv:%d, %d, %v\n", pv, j, a)
	return j
}

func main() {
	a := []int{2, 13, 7, 5, 9, 7, 11, 9, 10}
	quick(a, 0, len(a)-1)
}
