package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

//定义一个全局对象db
var db *sql.DB

//定义一个初始化数据库的函数
func initDB() (err error) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/golang_db?charset=utf8mb4&parseTime=True" //不会校验账号密码是否正确
	//注意!!!这里不要使用:=，我们是给全局变量赋值，然后在main函数中使用全局变量db
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	//尝试与数据库建立连接（校验dsn是否正确)
	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}

func insertData(name string, pwd string) {
	sqlStr := "insert into user_tb1(username, password) values (?, ?)"
	ret, err := db.Exec(sqlStr, name, pwd)
	if err != nil {
		fmt.Printf("insert failed! err : %v\n", err)
		return
	}
	theID, err := ret.LastInsertId()
	if err != nil {
		fmt.Printf("get last insert id failed! err : %v\n", err)
		return
	}
	fmt.Printf("get last insert id success, the id is : %v\n", theID)
}

func main() {

	// step1

	//db0, err := sql.Open("mysql", "root:123456@/golang_db")
	//if err != nil {
	//	panic(err)
	//}
	//
	//// See "Important settings" section.
	////最大连接时长
	//db0.SetConnMaxLifetime(time.Minute * 3)
	////最大连接数
	//db0.SetMaxOpenConns(10)
	////空闲连接数
	//db0.SetMaxIdleConns(10)
	//
	//fmt.Println(db0)

	//step2
	err1 := initDB()
	if err1 != nil {
		fmt.Printf("初始化失败！err : %v\n", err1)
		return
	} else {
		fmt.Println("初始化成功！")
	}

	insertData("wc", "wc222!")
}
