package main

import "fmt"

func main() {
	/**
	switch 语句不允许case表达式中存在结果值相等的自表达式，但是索引表达式却可以绕过编译器
	绕过方式对于类型判断是无效的，因为类型switch语句中case表达式的子表达式，都不许直接由类型字面量表示
	 */
	value2 :=[...]int8{0, 1, 2, 3, 4, 5, 6}
	switch value2[2] {
	case value2[0], value2[1], value2[2]:
		fmt.Println("0 or 1")
	case value2[2], value2[3], value2[4]:
		fmt.Println("2 or 3")

	case 4, 5:
		fmt.Println("4 or 5")
	}
}