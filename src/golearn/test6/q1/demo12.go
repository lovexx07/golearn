package main

import "fmt"

var container = []string{"zero", "one","two"}

func main() {
	container := map[int]string{0: "zero", 1:"one", 2:"tow"}

	//fmt.Printf("The element is 1 is %s.\n", container[1])

	// 方式一 获取类型  所有对象都是interface 的子类
	/**
	判断类型的基本方法是  interface{}(X).(类型)
	 */
	_, ok1 := interface{}(container).([]string)
	_, ok2 := interface{}(container).(map[int]string)

	if !(ok1 || ok2){
		fmt.Printf("Error : unsupported container type: %T\n", container)
		return
	}

	fmt.Printf("The element is %q. (container type :%T)\n", container[1], container)

	//使用switch 语句判断类型 interface{}(X).(type){ ....
	elem, err :=getElement(container)
	if err!= nil{
		fmt.Printf("Error: %s\n", err)
	}
	fmt.Printf("The element is %q. (container type: %T)\n", elem, container)
	}


func getElement(containerI interface{}) (elem string, err error){
	switch t := containerI.(type) {
	case []string:
		elem = t[1]
	case map[int]string:
		elem = t[1]
	default:
		err = fmt.Errorf("unsupported container type: %T", containerI)
	}
	return
}