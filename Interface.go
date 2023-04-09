package main

import (
	"fmt"
	"os"
)

type Inter interface {
	Ping()
	Pang()
}

type Anter interface {
	Inter
	String()
}

type St struct {
	Name string
}

func (St) Ping() {
	println("ping...")
}

func (*St) Pang() {
	println("pang...")
}

func main() {
	st := &St{"andes"}
	var i interface{} = st

	//判断i绑定的实例是否实现了接口类型Inter
	if o, ok := i.(Inter); ok {
		o.Ping()
		o.Pang()
	}

	if p, ok := i.(Anter); ok {
		//i没有实现接口Anter，所以执行不到这里
		p.String()
	}

	//判断i绑定的实例是否就是具体类型St
	if s, ok := i.(*St); ok {
		fmt.Printf("%s", s.Name) //andes
	}

	f, err := os.OpenFile("note.txt", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		println("error...")
	}
	defer f.Close()
}
