package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("hello world11111")
	println("hello world2")
	var a string = "NULL"
	var b int = 20
	var c, d = 0.5, "OK123"
	println(a, b, c, d)

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/hello", helloHandler)
	log.Fatal(http.ListenAndServe(":9999", nil))
}

func indexHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
}

func helloHandler(w http.ResponseWriter, req *http.Request) {
	//fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
	for k, v := range req.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
}
