package main

import (
	"golearn/test37/common"
	"fmt"
	"os"
	"errors"
	"runtime/pprof"
	"golearn/test37/common/op"
)

var (
	profileName = "cpuprofile.out"
)

func main() {
	f, err := common.CreateFile("", profileName)
	if err != nil{
		fmt.Printf("CPU profile creation error: %v\n", err)
	}
	defer f.Close()
	if err := startCPUProfile(f); err!=nil{
		fmt.Printf("CPU profile start error: %v\n", err)
		return
	}

	if err = common.Execute(op.CPUProfile, 10); err!= nil{
		fmt.Printf("excute error: %v\n", err)
		return
	}
	stopCPUProfile()

}

func startCPUProfile(f *os.File) error {
	if f == nil{
		return errors.New("nil file")
	}
	return pprof.StartCPUProfile(f)
}

func stopCPUProfile() {
	pprof.StopCPUProfile()
}