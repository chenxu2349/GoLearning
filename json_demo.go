package main

import (
	"encoding/json"
	"fmt"
)

// Person 序列化成json格式时需要所有字段首字母大写，不大写那个字段就会丢失
type Person struct {
	Name  string
	Age   int
	Email string
}

func main() {

	p := Person{
		Name:  "Andy",
		Age:   19,
		Email: "andy@gmail.com",
	}

	p_json, _ := json.Marshal(p)
	//print(p_json) //需要强转成string类型
	fmt.Println(string(p_json))

	jp := []byte(`{"Name":"Andy","Age":19,"Email":"andy@gmail.com"}`) //反引号
	var p2 Person
	json.Unmarshal(jp, &p2) //需要取地址
	fmt.Printf("p2 : %v\n", p2)

}
