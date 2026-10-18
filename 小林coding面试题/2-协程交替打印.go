package main

import (
	"fmt"
	"sync"
)

//三个goroutinue交替打印abc 10次

func main() {
	// 定义3个channel
	ch1 := make(chan struct{})
	ch2 := make(chan struct{})
	ch3 := make(chan struct{})
	var wg sync.WaitGroup
	wg.Add(3)
	// 打印a
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			<-ch1
			fmt.Println("a")
			ch2 <- struct{}{}
		}
		// 第10次的时候，打印c的goroutine写入了ch1
		// 为了防止阻塞，要消费一下ch1
		<-ch1
	}()
	// 打印b
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			<-ch2
			fmt.Println("b")
			ch3 <- struct{}{}
		}
	}()
	// 打印c
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			<-ch3
			fmt.Println("c")
			ch1 <- struct{}{}
			fmt.Println("第", i+1, "轮结束")
		}
	}()
	// 启动
	ch1 <- struct{}{}
	wg.Wait()
	close(ch1)
	close(ch2)
	close(ch3)
	fmt.Println("end")
}
