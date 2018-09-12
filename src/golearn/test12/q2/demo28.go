package main

import "fmt"

//普通数字是值类型，传递的是副本
func modifyArray(a [3]string) [3]string {
	a[1] = "X"
	fmt.Printf("指针1:%p\n", &a)
	return a
}

//切片类型 传递的是引用，任何数据操作都会改变原值
func modifySlice(a []string) []string {
	a[1] = "i"
	return a
}
//数组切片类型，修改切片导致原始数据变动，只修改数组只修改了副本
func modifyComplexArray(a [3][]string) [3][]string {
	a[1][1] = "s"    //切片 操作引用
	a[2] = []string{"o", "p", "q"}  //数组 操作副本
	return a
}

func main() {
	array1 := [3]string{"a", "b", "c"}
	fmt.Printf("The array: %v\n", array1)
	array2 := modifyArray(array1)
	fmt.Printf("The array: %v\n", array2)
	fmt.Printf("The array: %v\n", array1)
	fmt.Printf("指针2：%p\n", &array2)
	fmt.Println()

	slice1 := []string{"x", "y", "z"}
	fmt.Printf("The slice: %v\n", slice1)
	slice2 := modifySlice(slice1)
	fmt.Printf("The modifyed slice: %v\n", slice2)
	fmt.Printf("THe original slice: %v\n", slice1)
	fmt.Println()

	complexArray1 :=[3][]string{
		[]string{"d", "e", "f"},
		[]string{"g", "k", "h"},
		[]string{"x", "y", "z"},
	}
	fmt.Printf("The complex array: %v\n", complexArray1)
	complexArray2 := modifyComplexArray(complexArray1)
	fmt.Printf("The modify complex array: %v\n", complexArray2)
	fmt.Printf("The original complex array: %v\n", complexArray1)
}