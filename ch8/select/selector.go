package main

import (
	"fmt"
	"os"
	"time"
)

func main()  {
	keyboardin := make(chan []byte)
	// 从键盘读取数据
	go func() {
		input := make([]byte, 1)
		_, err := os.Stdin.Read(input) // read a single byte
		if err == nil {
			keyboardin <- input
		}
	}()
	tick := time.Tick(5 * time.Second)
	select {
	case x := <-keyboardin:
		fmt.Printf("key board input : %s\n", x)
	case x := <-tick:
		fmt.Println("time out : ", x)

	}
}
