package main

import (
	"fmt"
	"sort"
)

type Student struct {
	Age    uint32
	Weight float32
}

type StuSlice []Student

func (s StuSlice) Len() int {
	return len(s)
}

func (s StuSlice) Less(i, j int) bool {
	return s[i].Age < s[j].Age || (s[i].Age == s[j].Age && s[i].Weight <= s[j].Weight)
}

func (s StuSlice) Swap(i, j int)  {
	s[i], s[j] = s[j], s[i]
}

func main() {
	stuSlice := StuSlice{}
	stuSlice = append(stuSlice, Student{10, 40})
	stuSlice = append(stuSlice, Student{5, 40})
	sort.Sort(stuSlice)
	fmt.Println(stuSlice)
}
