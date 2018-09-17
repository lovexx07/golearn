package main

import (
	"fmt"
)

func main() {
	num :=10
	sign := make(chan struct{}, 20)

	for i:=0; i<num; i++{
		go func() {
			fmt.Println(i)
			sign <- struct{}{}
		}()
	}

	//方法一 让主go语句等待  输出接口随机 很有可能10 个10
	//time.Sleep(time.Millisecond* 500)

	//利用通道等待执行  输出乱
	for j:=0; j< num; j++{
		<-sign
	}
}