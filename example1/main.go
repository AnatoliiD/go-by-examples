package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

// Item is an awesome json model 🤘
type Item struct {
	Name   string `json:"name"`
	Number int    `json:"number"`
}

var items []Item

func init() {
	for i := 1; i <= 10; i++ {
		name := fmt.Sprintf("Awesome Item#%d", i)
		fmt.Println(name)
		items = append(items, Item{
			Name:   name,
			Number: i,
		})
	}
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("static")))
	http.HandleFunc("/api/items", func(w http.ResponseWriter, r *http.Request) {
		e := json.NewEncoder(w)
		w.Header().Add("Content-Type", "application/json")
		e.Encode(&items)
	})
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
