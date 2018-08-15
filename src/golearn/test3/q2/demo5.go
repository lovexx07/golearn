package main

import (
	"flag"
	"golearn/test3/q2/lib"
)
var name string

func init() {
	flag.StringVar(&name, "name", "everyon", "object")
}

func main() {
	flag.Parse()
	lib.Hello(name)
}