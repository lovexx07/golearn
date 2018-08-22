package main

import (
	"flag"
	"golearn/test3/q4/lib"
)

var name string

func init() {
	flag.StringVar(&name, "name", "everyone", "The greeting object")

}

func main() {
	flag.Parse()
	lib.Hello(name)
}
