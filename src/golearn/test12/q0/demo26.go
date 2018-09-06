package main

import "fmt"

//定义一个函数类型
type Printer func(content string)(n int, err error)

//定义一个函数
func printToStd(content string) (bytesNum int, err error) {
	return fmt.Println(content)
}

func main() {
	var p Printer   //定义一个函数类型p
	p = printToStd   //给p 指向一个具体方法实现
	p("something")
}