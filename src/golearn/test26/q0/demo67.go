package main

import (
	"sync/atomic"
	"time"
	"fmt"
	"sync"
)

func main() {
	coordinateWithWaitGroup()
}
//先统一add 再并发done, 最后wait
func coordinateWithWaitGroup(){
	total :=12
	stride := 3
	var num int32
	fmt.Printf("The number: %d [with sync.WaitGroup]\n", num)
	var wg sync.WaitGroup
	fmt.Println("Start loop...")
	for i := 1; i <= total; i++ {
		wg.Add(stride)					//add 创建3哥需要等待的goroutine数量
		for j := 0; j < stride; j++ {
			go addNum(&num, i+j, wg.Done)
		}
		wg.Wait()
	}
	fmt.Println("End")
}

func addNum(numP *int32, id int, deferFunc func())  {
	defer func() {
		deferFunc()
	}()

	for i := 0; ; i++ {
		currNum := atomic.LoadInt32(numP)
		newNum := currNum+1
		time.Sleep(time.Millisecond*200)
		if atomic.CompareAndSwapInt32(numP, currNum, newNum){
			fmt.Printf("The number: %d [%d-%d]\n", newNum, id, i)
			break
		}else{

		}
	}
}