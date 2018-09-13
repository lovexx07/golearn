package main

type Dog struct {
	name string
}

func New(name string) Dog{
	return Dog{name:name}
}

func (dog *Dog) SetName(name string){
	dog.name = name
}

func (dog Dog) Name() string {
	return dog.name
}

func main() {
	//New函数的调用结果是不可寻址的临时变量，无法进行取值操作。 不能取得地址
	//New("little pig").SetName("monster")

	//素日对字典字面量和字典变量索引表达式的结果是不可寻址的，但是这样的表达式却可以用在自增或者自减上
	map[string]int{"the": 0, "word": 0, "counter": 0}["word"]++
	map1 :=map[string]int{"the": 0, "word": 0, "counter": 0}
	map1["word"] ++


}