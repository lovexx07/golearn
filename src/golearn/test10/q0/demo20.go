package main

import "fmt"

func main() {
	//通道创建
	ch1 := make(chan int, 3)
	ch1 <- 2
	ch1 <- 1
	ch1 <- 3

	elem1 :=  <-ch1  //接收一个元素
	fmt.Printf("The first element received from channel ch1 : %v\n", elem1)
}
