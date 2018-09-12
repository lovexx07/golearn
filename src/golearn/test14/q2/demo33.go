package main

import (
	"fmt"
	"reflect"
)

type Pet interface {
	Name() string
	Category() string
}

type Dog struct {
	name  string
}

func (dog *Dog) SetName(name string) {
	dog.name = name
}

func (dog Dog) Name() string{
	return dog.name
}

func (dog Dog) Category() string{
	return "dog"
}

func main() {
	var dog1 *Dog
	fmt.Println("The first dog is nil")
	dog2 := dog1
	fmt.Println("The second dog is nil")
	var pet Pet = dog2  // 赋予了pet一个 *Dog类型的nil,使用iface包装，之后就不是nil
	if pet == nil{
		fmt.Println("The pet is nil")
	}else {
		fmt.Println("The pet is not nil")
	}
	fmt.Printf("The type of pet is %T.\n", pet)
	fmt.Printf("The pyte of pet is %s.\n", reflect.TypeOf(pet).String())
	fmt.Println()

	wrap := func(dog *Dog) Pet {
		if dog == nil{
			return nil
		}
		return dog
	}
	pet = wrap(dog2)
	if pet == nil{
		fmt.Println("The pet is nil")
	}else{
		fmt.Println("The pet is not nil")
	}

}