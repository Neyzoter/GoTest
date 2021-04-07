package main

import "fmt"

func main() {
	mp := map[string]int{}
	mp["s"] = 3
	fmt.Println(mp)
	v, ok := mp["s"]
	if ok {
		fmt.Println(v)
	}
	delete(mp, "s")
	fmt.Println(mp)
}
