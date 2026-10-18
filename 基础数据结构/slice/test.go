package main

import "fmt"

func main() {
	s := []int{1, 2, 3}
	fmt.Println(s[len(s)-1])
	s2 := make([]int, 10)
	fmt.Println(s2)
	var s3 []string
	s3 = append(s3, "hello")
	s3 = append(s3, "world")
	fmt.Println(s3)
}
