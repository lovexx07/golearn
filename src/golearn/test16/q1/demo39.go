package main

import (
	"fmt"
	"time"
)

func main() {
	num :=10
	sign := make(chan struct{}, 20) //chan struct{} 代表既不包含任何字段也不拥有任何方法的空结构体的结构体类型，占内存0，在Go程序中值存在1份

	for i:=0; i<num; i++{
		go func() {
			fmt.Println(i)
			sign <- struct{}{}
		}()
	}

	//方法一 让主go语句等待  输出接口随机 很有可能10 个10
	time.Sleep(time.Millisecond* 5000)

	//利用通道等待执行  输出乱
	//for j:=0; j< num; j++{
	//	<-sign
	//}
}