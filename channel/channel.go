package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

//打印 chan 里 数据
func printer(ch chan int) {
	for i := range ch {
		fmt.Printf("chan 测试 %d \n", i)
	}
	wg.Done()
}

// 启动
func main() {
	//创建一个 chan
	c := make(chan int)
	// 启动一个线程
	go printer(c)
	wg.Add(1)
	//循环 把数据赋值给  c
	for i := 1; i < 10; i++ {
		c <- i
	}
	close(c)
	wg.Wait()
}
