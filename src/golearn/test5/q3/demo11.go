package main

import (
	. "golearn/test5/q3/lib_go"
	"fmt"
)

var hello string = "hello main"
func main() {
	Hello()
	Hello()
}

func Hello(){
	fmt.Printf("THe word is %s", hello)
}