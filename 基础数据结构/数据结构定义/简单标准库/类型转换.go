package main

import (
	"fmt"
	"strconv"
)

func main() {
	// byte转数字
	s := "12345"            // s[0] 类型是byte
	num := s[0] - '0'       // 1
	num2 := int(s[0] - '0') // 1
	fmt.Println(num2)
	str := string(s[0]) // "1"
	c := 1
	b := byte(c + '0') // '1'  byte类型
	fmt.Printf("type:%T, value: %d\n", num, num)
	fmt.Printf("type:%T, value: %s\n", str, str)
	fmt.Printf("type:%T, value: %c\n", b, b)
	i := "a"
	it := int(i[0])
	fmt.Printf("type:%T, value: %d\n", it, it)
	// 字符串转数字
	num1, _ := strconv.Atoi("123")
	fmt.Printf("字符串转数字结果: type:%T, value: %d\n", num1, num1)
	//数字转字符串
	str1 := strconv.Itoa(123)
	fmt.Printf("数字转字符串结果: type:%T, value: %s\n", str1, str1)
}
