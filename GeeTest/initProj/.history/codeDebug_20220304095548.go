/*
 * @Author: griffith
 * @Date: 2022-03-03 14:11:30
 * @LastEditors: griffith
 * @LastEditTime: 2022-03-04 09:55:34
 * @Description: here填写简介
 */
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

	fmt.Println(reflect.TypeOf(user))
	traversalString()
	a := []int{1,3,5,8,7}
	myTest(8,a...)
}


/**
 * @description: 遍历中文字符串
 * @param {*}
 * @return {*}
 */
func traversalString() {
	s := "pprof.cn博客"
	for i := 0; i < len(s); i++ { //byte
		fmt.Printf("%v(%c) ", s[i], s[i])
	}
	fmt.Println()
	for _, r := range s { //rune
		fmt.Printf("%v(%c) ", r, r)
	}
	fmt.Println()
}

//找出两个元素之和为8的下标
func myTest(target int,a ...int) {
	for i:=0;i<len(a);i++{
		other := target - a[i]
		for j:=0;j<len(a);j++{
			if a[j] == other {
				fmt.Printf("(%d,%d)\n",i,j)
			}
		}
	}
}

