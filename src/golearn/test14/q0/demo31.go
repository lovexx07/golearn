package main

import "fmt"

//定义接口
type Pet interface {
	SetName(name string)
	Name() string
	Category() string
}


type Dog struct {
	name string
}

func (dog *Dog) SetName(name string) {
	dog.name = name
}

func (dog Dog) Name() string{
	return dog.name
}

func (dog Dog)Category() string{
	return "dog"
}

func main() {
	dog := Dog{"little pig"}
	_, ok := interface{}(dog).(Pet)  //只实现了接口的值方法，不能称dog实现了Pet
	fmt.Printf("Dog implements interface Pet: %v\n", ok) //false
	_,ok = interface{}(&dog).(Pet)  //实现了pet的所有方法，包括值方法与指针方法
	fmt.Printf("*Dog implements interface Pet: %v\n", ok) //true
	fmt.Println()

	var pet Pet = &dog   //&dog 为pet的动态类型，Pet是pet的静态类型；静态类型不变，动态类型随赋值改变
	fmt.Printf("This not is a %s, the name is %q.\n", pet.Category(), pet.Name())

}