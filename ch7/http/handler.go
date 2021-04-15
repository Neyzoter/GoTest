package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	db := Database{"shoes": 50, "socks": 5}
	log.Fatal(http.ListenAndServe("localhost:8000", db))
}

type Dollars float32

func (d Dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

type Database map[string]Dollars

func (db Database) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		_, err := fmt.Fprintf(w, "%s : %s\n", item, price)
		if err != nil {
			log.Fatal(err)
		}
	}
	switch req.URL.Path {
	case "/list":
		fmt.Println("req list")
	default: // 未找到该path
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such page: %s\n", req.URL)
	}
}