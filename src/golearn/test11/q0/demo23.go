package main

import (
	"fmt"
	"math/rand"
)

func main() {
	//只能发不能收的通道
	var uselessChan = make(chan<- int, 1)
	_ = uselessChan


	//只能收不能发的通道
	var anthoerUseLessChan = make(<-chan int, 1)
	_=anthoerUseLessChan

	fmt.Printf("The useless channel: %v, %v\n", uselessChan, anthoerUseLessChan)

	//只收不发通道应用
	intChan1 := make(chan int, 3)
	SendInt(intChan1)

	//只发不收通道应用
	intChan2 := getIntChan()
	for elem := range intChan2{
		fmt.Printf("The element in intChan2: %v\n", elem)
	}

	//通道类型应用
	_ = GetIntChan(getIntChan)
}

func SendInt(ch chan <- int)  {
	ch <- rand.Intn(1000)
}

func getIntChan() <-chan int{
	num :=5
	ch := make(chan int, num)
	for i := 0; i< num; i++{
		ch <- i
	}
	close(ch)
	return ch
}

//通道类型使用 接口里面做规范
type Notifier interface {
	SendInt(ch chan<- int)
}

//通道使用 方法
type GetIntChan func() <-chan int