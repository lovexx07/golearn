package main

import (
	"golearn/test37/common"
	"fmt"
	"runtime"
	"os"
	"errors"
	"runtime/pprof"
	"golearn/test37/common/op"
)

var (
	profileName = "memprofile.out"
	memProfileRate = 8
)


func main() {
	f, err := common.CreateFile("", profileName)
	if err != nil{
		fmt.Printf("memory profile creation error: %v\n", err)
		return
	}

	defer f.Close()
	startMemProfile()
	if err = common.Execute(op.MemProfile, 10); err != nil{
		fmt.Printf("execute error: %v\n", err)
		return
	}
	if err :=stopMemProfile(f); err != nil{
		fmt.Printf("memory profile stop error: %v\n", err)
		return
	}
}

func startMemProfile() {
	runtime.MemProfileRate = memProfileRate
}

func stopMemProfile(f *os.File) error {
	if f == nil{
		return errors.New("nil file")
	}
	return pprof.WriteHeapProfile(f)
}