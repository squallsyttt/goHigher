package main

import "fmt"

func main() {
	messages := make(chan string, 2)

	messages <- "buff1"
	messages <- "buff2"

	//缓冲区先进先出
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}