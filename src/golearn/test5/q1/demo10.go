package main

import "fmt"

var block = "package"

func main(){
	block := "function"
	{
		block := "inner"
		fmt.Printf("The blcok is %s.\n", block)
	}
	fmt.Printf("The block is %s.\n", block)
}