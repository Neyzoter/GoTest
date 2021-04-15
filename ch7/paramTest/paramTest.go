package main

import "fmt"

func arrayChange(v [4]int) {
	v[1] = 10
}
func arrayPtrChange(v *[4]int) {
	v[1] = 10
}

func sliceChange(v []int) {
	v[1] = 10
}

func mapChange(v map[string]int) {
	v["123"] = 123
}

func main() {
	// 数组只能通过指针修改
	v := [4]int{1, 2, 3, 4}
	arrayChange(v)
	fmt.Println(v)
	arrayPtrChange(&v)
	fmt.Println(v)

	// slice可以直接更改
	p := []int{1, 2, 3, 4}
	sliceChange(p)
	fmt.Println(p)

	// map可以直接更改
	m := map[string]int{}
	m["123"] = 111
	mapChange(m)
	fmt.Println(m)
}
