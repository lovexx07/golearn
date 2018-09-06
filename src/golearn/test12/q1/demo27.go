package main

import (
	"errors"
	"fmt"
)
//定义一个函数类型 传入2个int值 返回一个int
type operate func(x, y int) int

//定义一个函数
func calculate(x int, y int, op operate) (int, error) {
	if op == nil{  //判断类型
		return 0, errors.New("invalid operation")
	}
	return op(x, y), nil
}

type calculateFunc func(x int, y int)(int, error)

//返回一个方法类型
func genCalculator(op operate) calculateFunc {
	return func(x int, y int) (int, error) {
		if op == nil{
			return 0, errors.New("invalid operation")
		}
		return op(x, y), nil
	}
}

func main() {
	x, y :=12, 23
	op := func(x, y int) int {
		return x + y
	}

	result, err := calculate(x, y, op)
	fmt.Printf("The result: %d (error: %v)\n", result, err)

	result, err = calculate(x, y, nil)
	fmt.Printf("The result: %d (error: %v)\n", result, err)


	x, y = 56, 78
	add := genCalculator(op)
	result, err = add(x, y)
	fmt.Printf("The result :%d (error :%v)\n", result, err)
}