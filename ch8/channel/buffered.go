package main

import (
	"fmt"
	"time"
)

// 单个channel只能被读取一次
// 不同的goroutine总共只能读取到一次
func main()  {
	ch := make(chan int, 3) // buffered
	go readChan1(ch)
	go writeChan1(ch)
	go rangeChan1(ch)
	time.Sleep(time.Second * 20)
}

// ch为一个无缓冲的输入channel
func readChan1(ch <-chan int)  {
	// 只能关闭输出channel，而不能关闭输入channel
	//close(ch)
	for i := 1; i < 10; i++ {
		// 如果没有数据则会阻塞
		rec, ok := <-ch
		if ok { // 是否从channel成功接收到数据
			fmt.Println("read : ", rec)
		} else { // 通道已经关闭
			fmt.Println("channel closed")
		}
		//time.Sleep(time.Second * 1)
	}
}

// ch为一个无缓冲的输出channel
func writeChan1(ch chan<- int)  {
	for i := 1; i < 10; i++ {
		ch <- i
		// 如果数据未被读取，则会阻塞
		fmt.Println("write : ", i)
		time.Sleep(time.Second * 1)
	}
	close(ch)
}

// ch为一个无缓冲的输入channel
func rangeChan1(ch <-chan int)  {
	for x := range ch {
		fmt.Println("range read : ", x)
	}
}
