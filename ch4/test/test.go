package main

import "fmt"

func main() {
	v := [3]int{1, 2, 3}
	ptr := &v
	// error
	//*ptr = 9
	ptr[0] = 9
	print1(ptr)
	print2(ptr)
	print3(v)
	print4(v)
}

func print1(ptr *[3]int) {
	for i := range ptr {
		fmt.Println(i)
	}
}

func print2(ptr *[3]int) {
	for i, num := range ptr {
		fmt.Printf("%d : %d\n", i, num)
	}
}

func print3(v [3]int) {
	for i := range v {
		fmt.Println(i)
	}
}

func print4(v [3]int) {
	for i, num := range v {
		fmt.Printf("%d : %d\n", i, num)
	}
}
