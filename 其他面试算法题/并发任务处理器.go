package main

import (
	"fmt"
	"sync"
	"time"
)

/*
实现一个并发任务处理器:给定一个包含100个任务ID的列表，要求控制最大并发数为3，
模拟并发调用某个外部接口(如打印ID)
 */
func main() {
	taskIDs := make([]int, 100)
	for i := 0; i < 100; i++ {
		taskIDs[i] = i
	}
	// 控制最大并发数为3
	maxConcurrent := 3
	semaphore := make(chan struct{}, maxConcurrent)
	var wg sync.WaitGroup
	// 处理任务
	for _, id := range taskIDs {
		wg.Add(1)
		// 获取信号量：塞不进去就阻塞
		semaphore <- struct{}{}
		go func(taskID int) {
			defer wg.Done()
			defer func() { <-semaphore }() // 释放信号量

			// 模拟调用外部接口
			fmt.Printf("Processing task %d\n", taskID)
			time.Sleep(10 * time.Millisecond) // 模拟接口调用耗时
		}(id)
	}
	// 等待所有任务完成
	wg.Wait()
	close(semaphore) 
	fmt.Println("All tasks completed")
}
// close(semaphore) 可以删掉
// Go 的 GC 会自动回收没有任何引用指向的 channel，无论它是否被 close 过。
// 这里的 semaphore 在 main() 函数退出后就脱离作用域，没有引用了，GC 自会回收。
/*
close 真正的作用是通知接收方"不再有数据"，比如配合 range 遍历 channel：
for v := range ch {  // 只有 close(ch) 后，range 才会退出
    ...
}
*/