package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type dollars float64

func (d dollars) String() string {
	s := strconv.Itoa(int(d))
	return s + "$"
}

func main() {
	db := database{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/update", db.update)
	http.HandleFunc("/create", db.create)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request) {
	temp := template.Must(template.ParseFiles("index.html"))
	if err := temp.Execute(w, db); err != nil {
		log.Println(err)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}

func (db database) create(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	value := req.URL.Query().Get("value")

	if item == "" || value == "" {
		fmt.Fprintln(w, "create error")
		return
	}

	if _, ok := db[item]; ok {
		fmt.Fprintln(w, "exist item")
	} else {
		num, err := strconv.Atoi(value)

		if err != nil {
			fmt.Fprintln(w, "create error")
		} else {
			db[item] = dollars(num)
		}
	}
}

func (db database) update(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	value := req.URL.Query().Get("value")

	if item == "" || value == "" {
		fmt.Fprintln(w, "update error")
		return
	}

	if _, ok := db[item]; !ok {
		fmt.Fprintln(w, "not exist item")
	} else {
		num, err := strconv.Atoi(value)

		if err != nil {
			fmt.Fprintln(w, "update error")
		} else {
			db[item] = dollars(num)
		}
	}
}
