package main

import (
	"./impl"
	"fmt"
)

func main() {
	// GetName的方法接收着为(t T)，T或者*T都可以创建实现
	// Values                    Methods Receivers
	//-----------------------------------------------
	// T                        (t T)
	//*T                        (t T) and (t *T)
	var fruit impl.Fruit = &impl.Apple{Name: "apple"}
	//fruit = impl.Apple{Name: "apple"} // error
	fmt.Println(fruit)
}