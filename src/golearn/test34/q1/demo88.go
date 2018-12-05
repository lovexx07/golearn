package main

import (
	"path/filepath"
	"os"
	"fmt"
	"io/ioutil"
)

type flagDesc struct {
	flag int
	desc string
}
func main() {
	fileName1 := "something2.txt"
	filePath1 := filepath.Join(os.TempDir(), fileName1)
	fmt.Printf("The file path: %s \n", filePath1)
	fmt.Println()

	content0 := "OpenFile is the generalized open call"
	flagDescList := []flagDesc{
		{
			os.O_WRONLY | os.O_CREATE | os.O_TRUNC,
			"os.O_WRONLY | os.O_CREATE | os.O_TRUNC",
		},
		{
			os.O_WRONLY,
			"os.O_WRONLY",
		},
		{
			os.O_WRONLY|os.O_APPEND,
			"os.O_WRONLY|os.O_APPEND",
		},
	}
	for i, v := range flagDescList {
		fmt.Printf("Open the file with flag %s ...\n", v.desc)
		filela, err := os.OpenFile(filePath1, v.flag, 0666)
		if err != nil {
			fmt.Printf("error: %v\n", err)
			continue
		}
		fmt.Printf("The file descriptor: %d\n", filela.Fd())

		contents1 := fmt.Sprintf("[%d]: %s ", i+1, content0)
		fmt.Printf("Write %q to the file ...\n", contents1)
		n, err := filela.WriteString(contents1)
		if err != nil{
			fmt.Printf("error: %v\n", err)
			continue
		}
		fmt.Printf("The number of bytes written is %d.\n", n)

		filelb, err := os.Open(filePath1)
		fmt.Println("Read bytes from the file...")
		bytes, err:= ioutil.ReadAll(filelb)
		if err != nil{
			fmt.Printf("error: %v\n", err)
			continue
		}
		fmt.Printf("Read(%d): %q\n", len(bytes), bytes)
		fmt.Println()

	}


	fmt.Println("Try to create an existiong file with flag os.O_TRUNC ....")
	file2, err:= os.OpenFile(filePath1, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil{
		fmt.Printf("error: %v\n", err)
		return
	}
	fmt.Printf("The file desriptor: %d\n", file2.Fd())

	fmt.Println("Try to create an existing file with flag os.O_EXCL ...")
	_, err = os.OpenFile(filePath1, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0666)
	fmt.Printf("error: %v\n", err)

}