package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func showMessage(s string) {
	for i := 0; i < 10; i++ {
		fmt.Println(s)
		time.Sleep(time.Millisecond * 500)
	}
}

func responseSize(url string) {

	fmt.Println("step1: ", url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("step2: ", url)
	defer response.Body.Close()

	fmt.Println("step3: ", url)
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("step4: ", len(body))
}

func main() {
	//go showMessage("Jack...") // go关键字启动一个协程
	//showMessage("Tom...")
	// fmt.Println("main end...")
	go responseSize("https://www.duoke360.com")
	go responseSize("https://baidu.com")
	go responseSize("https://jd.com")
	time.Sleep(10 * time.Second)
}
