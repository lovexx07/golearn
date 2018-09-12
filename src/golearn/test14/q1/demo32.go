package main

import "fmt"

type Pet interface {
	Name() string
	Category() string
}

type Dog struct {
	name string
}

func (dog *Dog) SetName(name string){
	dog.name = name
}

func (dog Dog) Name() string{
	return dog.name
}

func (dog Dog) Category() string{
	return "dog"
}

func main() {
	dog := Dog{"little pig"}
	fmt.Printf("The dog's name is %q,\n", dog.Name())
	var pet Pet = dog
	dog.SetName("monster")
	fmt.Printf("The dog's name is %q.n", dog.Name())
	fmt.Printf("This pet is a %s, the name is %q,\n", pet.Category(), pet.Name())
	fmt.Println()

	dog1 := Dog{"little pig"}
	fmt.Printf("The name of first dog is %q.\n", dog1.Name())
	dog2 := dog1 //赋值用的是 副本
	fmt.Printf("THe name of second dog is %q.\n", dog2.Name())
	dog1.name = "monster"
	fmt.Printf("The name of first dog is %q.\n", dog1.Name())
	fmt.Printf("The name of second dog is %q.\n", dog2.Name())
	fmt.Println()

	dog = Dog{"little pig"}
	fmt.Printf("The dog's name is %q.\n", dog.Name())
	pet = &dog  //给一个接口变量赋值时，该变量的动态类型会与它的值一起被存储在一个专用的数据结构中。包含了dog值的副本
	dog.SetName("monster")
	fmt.Printf("The dog's name is %q.\n", dog.Name())
	fmt.Printf("This pet is a %s, the name is %q.\n", pet.Category(), pet.Name())

}