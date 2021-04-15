package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	db := database{"shoes": 50, "socks": 5}
	mux := http.NewServeMux()
	mux.Handle("/list", http.HandlerFunc(db.list))
	mux.Handle("/shoes", http.HandlerFunc(db.shoes))
	mux.HandleFunc("/socks", db.socks)
	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}

type dollars float32

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

type database map[string]dollars


func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		_, err := fmt.Fprintf(w, "%s : %s\n", item, price)
		if err != nil {
			log.Fatal(err)
		}
	}
}
func (db database) shoes(w http.ResponseWriter, req *http.Request) {
	_, err := fmt.Fprintf(w, "%s : %s\n", "shoes", db["shoes"])
	if err != nil {
		log.Fatal(err)
	}
}
func (db database) socks(w http.ResponseWriter, req *http.Request) {
	_, err := fmt.Fprintf(w, "%s : %s\n", "socks", db["socks"])
	if err != nil {
		log.Fatal(err)
	}
}