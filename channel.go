package main

import (
	"fmt"
	"math/rand"
	"time"
)

var values = make(chan int)

//var chan1 = make(chan string 10)  //有缓冲通道，长度为10
//var chan2 = make(chan float64)  //无缓冲通道

//往通道里写随机数
func send() {
	rand.Seed(time.Now().UnixNano())
	value := rand.Intn(10)
	fmt.Printf("send : %v\n", value)
	time.Sleep(3 * time.Second)
	values <- value
}

func main() {
	//从通道接收值
	defer close(values)
	go send()
	fmt.Println("wait...")
	value := <-values
	fmt.Printf("receive : %v\n", value)
	fmt.Println("end...")
}
