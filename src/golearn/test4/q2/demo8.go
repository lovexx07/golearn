package main

import (
	"flag"
	"fmt"
)

func main() {
	var name = getTheFlag()
	flag.Parse()
	fmt.Printf("Hello, %v\n", *name)
}

func getTheFlag() *string {
	return flag.String("name", "everyone8", "the greeting object")
}
