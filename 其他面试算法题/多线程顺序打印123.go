package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	ch1 := make(chan struct{})
	ch2 := make(chan struct{})
	ch3 := make(chan struct{})
	wg.Add(3)
	go func() {
		defer wg.Done()
		<-ch1
		fmt.Println("1")
		ch2 <- struct{}{}
	}()
	go func() {
		defer wg.Done()
		<-ch2
		fmt.Println("2")
		ch3 <- struct{}{}
	}()
	go func() {
		defer wg.Done()
		<-ch3
		fmt.Println("3")
	}()
	// start
	ch1 <- struct{}{}
	wg.Wait()
	fmt.Println("done")
}
