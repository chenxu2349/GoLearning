package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	basicLearning()
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

func basicLearning() {
	fmt.Println("hello world11111")
	println("hello world2")
	var a string = "NULL"
	var b int = 20
	var c = "OK123"
	var d, e, f, g, h = 0.5, 1.e+0, 6.6748e-11, .1314e+5, 1e6 // 浮点型
	var ch rune = 'a'
	aa, bb := 1, "中国" // 直接赋值
	println(a, b, c, d, e, f, g, h, ch)
	println(aa, bb)
}
