package main

import (
	"flag"
	"fmt"
)

func main() {
	// var name string
	// flag.StringVar(&name, "name", "everyone", "The greeting object")

	var name = *flag.String("name", "everyone", "The greeting object")

	// name := *flag.String("name", "everyone1", "The greeting object")

	flag.Parse()
	fmt.Printf("name is %v\n", name)
}
