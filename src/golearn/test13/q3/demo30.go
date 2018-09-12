package main

import "fmt"

type Cat struct {
	name 			string
	scientificName  string
	category		string
}

func New(name, scientificName, category string) Cat {
	return Cat{
		name:	name,
		scientificName: scientificName,
		category: category,
	}
}

//指针类型  改变的结构体的值
func (cat *Cat) SetName(name string)  {
	cat.name = name
}

//值传递  改变的只是副本，不会影响外部
func (cat Cat) SetNameOfCopy(name string) {
	cat.name = name
}

func (cat Cat) Name() string  {
	return cat.name
}

func (cat Cat) ScientificName() string {
	return cat.scientificName
}

func (cat Cat) Category() string {
	return cat.category
}

func (cat Cat) String() string  {
	return fmt.Sprintf("%s (category: %s, name: %q)", cat.scientificName, cat.category, cat.name)
}

func main() {
	cat := New("little pig0", "American shorthair", "cat")
	cat.SetName("monster")
	fmt.Printf("The cat: %s\n", cat)

	cat.SetNameOfCopy("little pig1")
	fmt.Printf("The cat: %s\n", cat)

	type Pet interface {
		SetName(name string)
		Name() string
		Category() string
		ScientificName() string
	}

	_, ok := interface{}(cat).(Pet)
	fmt.Printf("Cat implements interface Pet: %v\n", ok)
	_, ok = interface{}(&cat).(Pet)
	fmt.Printf("*Cat implements interface Pet: %v\n", ok)

}