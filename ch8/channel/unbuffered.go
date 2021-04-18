package main

import (
	"fmt"
	"time"
)

func main()  {
	ch := make(chan int)
	defer close(ch)
	go readChan(ch)
	go writeChan(ch)
	time.Sleep(time.Second * 20)
}

// ch为一个无缓冲的channel
func readChan(ch chan int)  {
	for i := 1; i < 10; i++ {
		// 如果没有数据则会阻塞
		fmt.Println("read : ", <-ch)
		//time.Sleep(time.Second * 1)
	}
}

// ch为一个无缓冲的channel
func writeChan(ch chan int)  {
	for i := 1; i < 10; i++ {
		ch <- i
		// 如果数据未被读取，则会阻塞
		fmt.Println("write : ", i)
		time.Sleep(time.Second * 1)
	}
}
