package main

import (
	"fmt"
)

type Employee struct {
	Age int
	Weight float32
	// 匿名成员
	Birth
}

type Birth struct {
	Year int
	Month int
	Day int
}

func main() {
	employee := Employee{}
	fmt.Println(employee)

	employee1 := Employee{Age: 20, Weight: 80}
	fmt.Println(employee1)

	employee2 := Employee{20, 80, Birth{2000, 3, 30}}
	fmt.Println(employee2)
	// 可以讲Birth省略
	employee2.Year = 2022
	// #表示用和Go语言类似的语法打印值
	fmt.Printf("%#v", employee2)
}
