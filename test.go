package main

import (
	"flag"
	"fmt"
)

var a int

func init() {
	flag.IntVar(&a, "a", 123, "输入数字")
}

func main() {
	flag.Parse()
	fmt.Println("hello world ", a)
}
