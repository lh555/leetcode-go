package main

// 普通变量定义
a := 10
// 指针变量定义
b := &a
// 结构体变量定义
type Person struct {
	Name string
	Age  int
}
// 数组变量定义
c := [3]int{1, 2, 3}
d := &c
e := &c
// 切片变量定义
f := []int{1, 2, 3}
g := &f
h := &f
// 映射变量定义
i := map[string]int{"a": 1, "b": 2}
j := &i
k := &i
// 通道变量定义
l := make(chan int)
m := &l
n := &l
// 函数变量定义
o := func(a, b int) int {
	return a + b
}
p := &o
q := &o
// 接口变量定义
type Animal interface {
	Speak() string
}
