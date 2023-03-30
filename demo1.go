package main

import (
	"fmt"
)

func main() {
	basicLearning()
	//http.HandleFunc("/", indexHandler)
	//http.HandleFunc("/hello", helloHandler)
	//log.Fatal(http.ListenAndServe(":9999", nil))
}

//func indexHandler(w http.ResponseWriter, req *http.Request) {
//	fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
//}
//
//func helloHandler(w http.ResponseWriter, req *http.Request) {
//	//fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
//	for k, v := range req.Header {
//		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
//	}
//}

func basicLearning() {
	fmt.Println("hello world-11111")
	println("hello world-22222")
	var a string = "NULL"
	var b int = 20
	var c = "OK1234567"
	var d, e, f, g, h = 0.5, 1.e+0, 6.6748e-11, .1314e+5, 1e6 // 浮点型
	var ch rune = 'a'
	aa, bb := 1, "中国" // 直接赋值
	println(a, b, c, d, e, f, g, h, ch)
	println(aa, bb)
	subString := c[0:5] //字符串切片
	print(subString)
}

func basic2() {
	//复合数据类型：指针，数组，切片，字典（map），通道，结构，接口

	//pointer
	a := 11
	p := &a
	println(p, &a)

	type User struct {
		name string
		age  int
	}
	Jack := User{
		name: "Jack",
		age:  18,
	}
	q := &Jack
	println("name: " + q.name)

	//array
	var arr0 [2]int
	arr1 := [3]int{1, 4, 7}
	arr2 := [...]float64{7.0, 8.1, 9.5, 12.9}
	arr2Len := len(arr2)
	arr3 := [5]int{1: 3, 2: 6, 3: 1} //通过索引初始化值
	for i, v := range arr3 {
		println(i, v)
	}

	//slice:切片内部三个元素：指针，切片元素数量，底层数组容量
	var array = [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	s1 := array[0:4]
	s2 := array[:4]
	s3 := array[3:]
	s4 := make([]int, 10) //make([]int, len, cap)
	s5 := make([]int, 10, 15)

	s1 = append(s1, 100) //切片支持的操作：len(a), cap(a), a = append(a, 1), append(a, c...), copy(b, a)

	str := "hello, 世界！" // 字符串和切片相互转换
	a1 := []byte(str)
	b1 := []rune(str)

	//map[K]T

}

func sum(a, b int) *int {
	sum := a + b
	return &sum
}
