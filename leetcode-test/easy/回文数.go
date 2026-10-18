package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(x int) bool {
	s := strconv.Itoa(x) // 不能使用string(x)。Go 中 string(整数) 是把整数当作 Unicode 码点转成字符，而不是转成十进制字符串
	l := len(s)
	for i := 0; i < l/2; i++ {
		if s[i] != s[l-1-i] {
			return false
		}
	}
	return true
}

func main() {
	x := 10
	res := isPalindrome(x)
	fmt.Println("res:", res)
}
