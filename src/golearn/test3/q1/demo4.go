package q1

import "flag"

var name string

func init() {
	flag.StringVar(&name, "name", "evenyone", "the greeting object")

}

func main() {
	flag.Parse()
	hello(name)
}