package main

import (
	json2 "encoding/json"
	"fmt"
)

type Movie struct {
	Name string
	Year int `json:"released"`  // json编码后编程released
}

func main() {
	movie := Movie{Name: "Big", Year: 2021}
	fmt.Println(movie)
	json, err := json2.Marshal(movie)
	if err == nil {
		fmt.Printf("%s\n", json)
	}

	json, err = json2.MarshalIndent(movie, "", "  ")
	if err == nil {
		fmt.Printf("%s\n", json)
	} else {
		fmt.Println("marshal error")
	}

	err = json2.Unmarshal(json, &movie)
	if err == nil {
		fmt.Println(movie)
	} else {
		fmt.Println("unmarshal error")
	}
}
