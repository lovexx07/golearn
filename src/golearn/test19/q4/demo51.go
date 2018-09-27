package main

import "fmt"

func main() {
	// 多个defer 执行顺序，按照编译好的顺序从后往前执行
	defer fmt.Println("first defer")

	for i :=0 ; i<3; i++{
		defer fmt.Printf("defer in for [%d]\n", i)
	}
	defer fmt.Println("last defer")
}