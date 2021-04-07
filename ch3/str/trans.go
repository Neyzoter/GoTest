package main

import (
	"fmt"
	"strconv"
)

func main()  {
	// itoa将123转化为"123"
	fmt.Println(strconv.Itoa(123))
	intVal, err := strconv.Atoi("123")
	if err != nil {
		fmt.Println(intVal)
	} else {
		fmt.Println("atoi err")
	}

	// string可以将字符数值转化为字符，65 = 'A'
	fmt.Println(string(65))
}
