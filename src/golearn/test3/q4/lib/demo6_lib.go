package lib

import (
	internal "golearn/test3/q4/lib/internal"
	"os"
)

func Hello(name string) {
	internal.Hello(os.Stdout, name)
}
