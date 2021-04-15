package main

import "fmt"

// 创建匿名函数
func createNoname() func() int {
	x := 1
	return func() int {
		x = x + 1
		return x
	}
}

func main() {
	f := createNoname()
	fmt.Println(f())
	fmt.Println(f())
}
