package main

import "fmt"

func main() {
	numbers1 :=[]int{1, 2, 3, 4, 5, 6}

	/**
	range 表达式只会在for语句开始执行时被求值一次，无论后面会有多少次迭代
	range 表达式的求值结果会被复制，被迭代的兑现时range表达式结果值的副本而不是原值
	 */
	for i:= range numbers1{
		if i == 3{
			numbers1[i] |= i
		}
	}

	fmt.Println(numbers1)
	fmt.Println()

	numbers2 :=[...]int{1, 2, 3, 4, 5, 6}
	maxIndex2 := len(numbers2)-1

	for i, e:=range numbers2{  //拿副本进行迭代
		if i==maxIndex2{
			numbers2[0] +=e
		}else {
			numbers2[i+1] +=e
		}
	}
	fmt.Println(numbers2)
	fmt.Println()

	numbers3 :=[]int{1, 2, 3, 4, 5, 6}  //切片，操作会影响到range,因为直接操作的元数据
	maxIndex3 := len(numbers3)-1
	for i, e := range numbers3{
		if i== maxIndex3{
			numbers3[0] +=e
		}else{
			numbers3[i+1] = e
		}
	}
	fmt.Println(numbers3)
}