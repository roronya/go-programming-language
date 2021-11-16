package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

var mu sync.Mutex

func main() {
	db := database{"shoes": 50, "socks": 5}
	mux := http.NewServeMux()
	mux.HandleFunc("/list", db.list)
	mux.HandleFunc("/price", db.price)
	mux.HandleFunc("/update", db.update)
	mux.HandleFunc("/delete", db.remove)
	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}

type dollars float64

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}

func (db database) update(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	strPrice := req.URL.Query().Get("price")
	if item == "" || strPrice == "" {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}
	floatPrice, err := strconv.ParseFloat(strPrice, 64)
	if err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}
	mu.Lock()
	defer mu.Unlock()
	price := dollars(floatPrice)
	db[item] = price
	fmt.Fprintf(w, "%s: %s\n", item, price)
}

func (db database) remove(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if item == "" {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}
	mu.Lock()
	defer mu.Unlock()
	delete(db, item)
	fmt.Fprintf(w, "remove %s", item)
}
