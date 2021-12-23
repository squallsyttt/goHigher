package main

import "fmt"

type CalculateType func(a, b int) int

func add(a, b int) int {
	return a + b
}

func mul(a, b int) int {
	return a * b
}

func Calculate(a, b int, f CalculateType) int {
	return f(a, b)
}

func (f CalculateType) CalculateObj(a, b int) int {
	return f(a, b)
}

func main() {
	a, b := 2, 3
	fmt.Println(Calculate(a, b, add))
	fmt.Println(Calculate(a, b, mul))

	//fmt.Println(CalculateObj(a, b))
}

//func CalculateObj(a int, b int) interface{} {
//	return
//}

//func CalculateObj(a int, b int) interface{} {
//}
