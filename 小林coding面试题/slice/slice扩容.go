package main

import "fmt"

func main() {
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := slice[2:5] // s1 = [2,3,4], len=3, cap=8
	s2 := s1[2:6:7]  // high=6	到 s1 的第 6 个元素（不含）,max=7容量上限到 s1 的第 7 个元素（不含）	cap = 7-2 = 5
	s3 := s1[2:6]
	s2 = append(s2, 100)
	s2 = append(s2, 200)

	s1[2] = 20

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(slice)
	fmt.Println("S3=", cap(s3), "S2=", cap(s2))
}
