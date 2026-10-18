package main

import (
	"fmt"
	"time"
)

// a <-chan int 表示 a 是一个只读 channel（receive-only channel），元素类型为 int。
// a chan int双向 channel，可读可写
func goroutineA(a <-chan int) {
	val := <-a
	fmt.Println("goroutine A received data: ", val)
	return
}

func goroutineB(b <-chan int) {
	val := <-b
	fmt.Println("goroutine B received data: ", val)
	return
}

func main() {
	ch := make(chan int)
	go goroutineA(ch)
	go goroutineB(ch)
	ch <- 3
	time.Sleep(time.Second)

	//ch1 := make(chan struct{})
}
