package main

import (
	"fmt"
	"reflect"
	"strconv"
)

type User struct {
	Name string
	Age  int
}

func (User) GetUser(user User) string {
	return user.Name + " " + strconv.Itoa(user.Age)
}

func main() {
	user := User{"xiaoming", 13}
	//Go默认形式
	fmt.Printf("%v", user)
	fmt.Println(111)
	//类型+值对象
	fmt.Printf("%#v", user)
	fmt.Println(222)
	//输出字段名和字段值形式
	fmt.Printf("%+v", user)
	fmt.Println(333)
	//值类型的Go语法表示形式
	fmt.Printf("%T", user)
	fmt.Println(444)
	fmt.Printf("%%")

	fmt.Println(reflect.TypeOf(user))
}
