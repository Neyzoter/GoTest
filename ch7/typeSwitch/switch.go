package main

import "fmt"

type apple struct {
	Name string
	Weight float32
}

type purple struct {
	Name string
	Weight float32
}

func main()  {
	a := interface{}(apple{"apple", 32.0})
	switcher(a)
	ifSwitcher(a)
}

func switcher(a interface{})  {
	switch a.(type) {
	case apple:
		fmt.Println("apple's type is apple")
	case purple:
		fmt.Println("apple's type is apple")
	default:
		fmt.Println("Un-know type")
	}
}

func ifSwitcher(a interface{})  {
	if _, ok := a.(apple); ok {
		fmt.Println("apple's type is apple")
	} else if _, ok := a.(purple); ok {
		fmt.Println("apple's type is apple")
	} else {
		fmt.Println("Un-know type")
	}
}
