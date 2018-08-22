package main

import "fmt"

var container = []string{"zero", "one", "two","three"}

func main(){
	container := map[int]string{0: "zero", 1:"one1", 2:"tow"}
	fmt.Printf("The element is %q.\n", container[1])
}