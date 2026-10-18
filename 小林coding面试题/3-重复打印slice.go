package main

import (
	"fmt"
	"sync"
)

// 用不超过10个goroutine不重复的打印slice中的100个元素
// 容量为10的有缓冲channel实现
// 每次启动10个，累计启动100个goroutine,且无序打印
func main() {
	var wg sync.WaitGroup
	// 创建切片
	ss := make([]int, 100)
	for i := 0; i < 100; i++ {
		ss[i] = i
	}
	ch := make(chan struct{}, 10)
	for i := 0; i < 100; i++ {
		wg.Add(1)
		ch <- struct{}{} // 如果channel满了，就阻塞在这，等待goroutine打印完
		// 写10个就阻塞了，此时goroutine中打印
		go func(idx int) {
			defer wg.Done()
			fmt.Printf("val: %d \n", ss[idx])
			// 打印结束，从缓冲channel中删除一个
			<-ch
		}(i)

	}
	wg.Wait()
	// 关闭channel
	close(ch)
	fmt.Println("end")
}

// 用不超过10个goroutine不重复的打印slice中的100个元素
// 创建10个无缓冲channel和10个goroutine
// 固定10个goroutine,且顺序打印
func test9() {
	var wg sync.WaitGroup
	// 创建切片
	ss := make([]int, 100)
	for i := 0; i < 100; i++ {
		ss[i] = i
	}
	// 创建channel和goroutine
	hashMap := make(map[int]chan int)
	sort := make(chan struct{})
	for i := 0; i < 10; i++ {
		hashMap[i] = make(chan int)
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			for val := range hashMap[idx] {
				fmt.Printf("go id: %d, val: %d \n", idx, val)
				sort <- struct{}{}
			}
		}(i)
	}
	// 循环切片，对10取模，找到对应channel的key，写入值
	for _, v := range ss {
		id := v % 10
		hashMap[id] <- v
		// 有序
		<-sort
	}
	// 循环结束关闭channel,删除map的key
	for k, _ := range hashMap {
		close(hashMap[k])
		delete(hashMap, k)
	}
	wg.Wait()
	close(sort)
	fmt.Println("end")
}
