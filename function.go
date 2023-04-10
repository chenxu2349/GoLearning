package main

import "fmt"

func sayHello(name string) {
	fmt.Printf("Hello,%s", name)
	fmt.Println()
}

func test(name string, f func(string)) {
	f(name)
}

func sum(a int, b int) int {
	return a + b
}

func sub(a int, b int) int {
	return a - b
}

func calculator(op string) func(int, int) int {
	switch op {
	case "+":
		return sum
	case "-":
		return sub
	default:
		return nil
	}
}

func main() {
	test("Jack", sayHello)
	ff := calculator("+")
	fmt.Println(ff(2, 3))

	//匿名函数
	x := func(a int, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}(1, 3) //自己调用自己，如果这里不传参数那么x就是一个函数变量
	print(x)

	//指针
	var ip *int
	fmt.Printf("ip : %v\n", ip)
	fmt.Printf("ip : %T\n", ip)

	a := 10
	ip = &a
	fmt.Printf("ip : %v\n", ip)
	fmt.Printf("ip : %v\n", *ip)
}
