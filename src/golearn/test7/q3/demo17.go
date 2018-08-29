package main

import "fmt"

func main() {
	a1 := [7]int{1, 2, 3, 4, 5, 6, 7}
	fmt.Printf("a1: %v (len: %d, cap: %d)\n", a1, len(a1), cap(a1))

	//切片的容量是切的第一个位置到元数据的最后一位
	s9 :=a1[1: 4]
	fmt.Printf("s9: %v (len: %d, cap: %d)\n", s9, len(s9), cap(s9))
	fmt.Println()

	//超过原来长度是，重新生产新的，不在使用原来的
	for i := 1; i <= 5; i++ {
		s9 = append(s9, i)
		fmt.Printf("s9: %v (len: %d, cap: %d)\n", s9, len(s9), cap(s9))
	}

	fmt.Printf("a1: %v (len: %d, cap: %d)\n", a1, len(a1), cap(a1))
}