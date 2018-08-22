package lib_go

import "fmt"

var hello string = "hello"
var word string ="word"

func Hello() {
	fmt.Printf("The word is %s %s.\n", hello, word)
}