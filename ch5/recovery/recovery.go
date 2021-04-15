package main

import (
	"fmt"
)

type zeroDivide struct {
}

func main()  {
	compute(0)
}

func compute(val int) {
	defer func() {
		switch p := recover(); p {
		case nil:
			fmt.Println("No Panic")
		case zeroDivide{}:
			fmt.Println("Recoverred a panic", p)
		default:
			fmt.Println("Pass a panic : ", p)
			panic(p)
		}
	}()
	if val == 0 {
		panic(zeroDivide{})
	}
	fmt.Println(1.0 / val)
}
