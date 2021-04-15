package main

import (
	"fmt"
	"os"
	"runtime"
)

func main()  {
	deferFunc()
}

// defer语句执行顺序和声明顺序相反
func deferFunc()  {
	// 打印堆栈信息
	defer printStack()
	trace("first line")
	defer trace("1")
	defer trace("2")
	trace("last line")
}

func trace(info string)  {
	fmt.Println(info)
}
func printStack() {
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	os.Stdout.Write(buf[:n])
}
