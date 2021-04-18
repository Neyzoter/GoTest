package main

import "fmt"

type Strawberry struct {
	Weight float32
	Name string
}

type Purple struct {
	Weight float32
	Name string
}

// 通过断言来判断类型
// 针对不同的类型进行相应的处理
func main()  {
	strawberry := interface{}(Strawberry{10.0, "sb"})
	if _, ok := strawberry.(Strawberry); ok {
		fmt.Println("Strawberry : ", strawberry)
	}
	if _, ok := strawberry.(Purple); ok {
		fmt.Println("Purple : ", strawberry)
	}
}
