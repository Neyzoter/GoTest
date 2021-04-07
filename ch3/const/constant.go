package main

import "fmt"

type Weekday int

const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

const (
	One int = 1
	Two
)

func main()  {
	fmt.Println(Saturday)
	fmt.Printf("Two is not %d\n", Two)
}
