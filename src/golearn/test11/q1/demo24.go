package main

import (
	"math/rand"
	"fmt"
	"time"
)

func main() {
	example1()
	example2()
}

func example1() {
	//准备通道
	intChannles := [3]chan int{
		make(chan int, 1),
		make(chan int, 1),
		make(chan int, 1),
	}

	//随机选择一个通道，并向它发送元素值
	index := rand.Intn(3)
	fmt.Printf("The index: %d\n", index)

	intChannles[index] <- index

	//哪一个通道中有可取的元素值，哪个对应的分支就会执行
	select {
	case <-intChannles[0]:
		fmt.Printf("The first candidate case is selected")
	case <-intChannles[1]:
		fmt.Printf("The second condidate case is selected")
	case elem:= <-intChannles[2]:
		fmt.Printf("The third condidate case is selected, the element is %d\n", elem)
	default:
		fmt.Println("No condidate case is selected")
	}
}

func example2() {
	intChan := make(chan int, 1)
	//一秒后关闭通道
	time.AfterFunc(time.Second, func() {
		close(intChan)
	})
	//通道没有取到值，阻塞
	select {
	case _, ok := <-intChan:
		if !ok{
			fmt.Printf("The condidate case is closed")
			break
		}
		fmt.Println("The condidate case is selected.")
	}
}