package main

import "fmt"

func main() {
	//defer先进后出
	defer func() {
		println("first")
	}()

	defer func() {
		println("second")
	}()

	println("demo2")

	//类型方法
	var a MyInt = 10
	var b MyInt = 10

	//int支持的运算，新定义类型同样支持
	c := a + b
	d := a * b
	fmt.Printf("%d\n", c)
	fmt.Printf("%d\n", d)
}

type Map map[string]string

func (m Map) Print() {
	for _, i := range m {
		println(i)
	}
}

type MyInt int
