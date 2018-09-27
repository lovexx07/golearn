package main

import (
	"fmt"
	"errors"
)

func main() {
	/**
	panic 意思是抛出一个异常， 和python的raise用法类似
	recover是捕获异常，和python的except用法类似
	defer会延迟函数到其他函数之后完之后再执行，后面跟的是函数
	golang 的错误处理流程：当一个函数在执行过程中出现了异常或遇到
	panic()，正常语句就会立即终止，然后执行 defer 语句，再报告异
	常信息，最后退出 goroutine。如果在 defer 中使用了 recover()
	函数,则会捕获错误信息，使该错误信息终止报告。
	 */
	fmt.Println("Enter function main.")

	defer func() {
		fmt.Println("Enter defer function.")

		if p := recover();p !=nil{
			fmt.Printf("1panic: %s\n", p)
		}

		fmt.Println("Exit defer function.")
	}()

	//recover 函数错误用法
	fmt.Printf("no panic: %v\n", recover())
	//引发一个panic
	panic(errors.New("someting worng"))

	p := recover()
	fmt.Printf("panic: %s\n", p)

	fmt.Println("Exit function main.")
}