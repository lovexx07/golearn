package main

import (
	"sync/atomic"
	"fmt"
	"errors"
	"io"
	"os"
	"reflect"
)

func main() {
	var box atomic.Value
	fmt.Println("Copy box to box2.")
	box2 := box //原子值在真正使用前可以被复制

	v1 := [...]int32{1, 2, 3}//定义数组
	fmt.Printf("Sore %v to box.\n", v1)
	box.Store(v1)
	fmt.Printf("The value load from box is %v.\n", box.Load())
	fmt.Printf("The value load from box2 is %v.\n", box2.Load())
	fmt.Println()

	v2 := "123"
	fmt.Printf("Store %q to box2.\n", v2)
	box2.Store(v2)
	fmt.Printf("The value load box is %v.\n", box.Load())
	fmt.Printf("The value load box2 is %v.\n", box2.Load())
	fmt.Println()

	var box4 atomic.Value
	v4 := errors.New("something wrong")
	fmt.Printf("Store an error with message %q to box4.\n", v4)
	box4.Store(v4)
	v41 := io.EOF
	fmt.Println("Store a value of the some type to box4.")
	box4.Store(v41)
	v42, ok := interface{}(&os.PathError{}).(error)
	if ok{
		fmt.Printf("Store a value of type %T that implements errors interface to box4.\n", v42)
		//box4.Store(v42)  //这里会报错 panic , 提示存储值类型不一致
	}
	fmt.Println()

	box5, err := NewAtomicValue(v4)
	if err != nil{
		fmt.Printf("errors : %s\n", err)
	}
	fmt.Printf("The legal type in box5 is %s.\n", box5.TypeOfValue())
	fmt.Println("Store a value of the same to box5.")
	err = box5.Store(v41)
	if err != nil{
		fmt.Printf("error: %s\n", err)
	}
	fmt.Printf("Store a value of type %T that implements error interface to box5.\n", v42)
	if err != nil {
		fmt.Printf("error: %s\n", err)
	}
	fmt.Println()


	var box6 atomic.Value
	v6 := []int{1, 2, 3}
	fmt.Printf("Store %v to box6.\n", v6)
	box6.Store(v6)

	v6[1] = 4//不安全
	fmt.Printf("The value load from box6 is %v.\n", box6.Load())
	v6 = []int{1, 2, 3}
	store := func(v []int) {
		replice := make([]int, len(v))
		copy(replice, v)
		box6.Store(replice)
	}

	fmt.Printf("Store %v to box6.\n", v6)
	store(v6)
	v6[2] = 5
	fmt.Printf("The value load from box6 is %v.\n", box6.Load())
}

type atomicValue struct {
	v atomic.Value
	t reflect.Type
}


func NewAtomicValue(example interface{}) (*atomicValue, error)  {
	if example == nil{
		return nil, errors.New("atomic value: nil example")
	}
	return &atomicValue{
		t:reflect.TypeOf(example),
	}, nil
}

func (av *atomicValue) Store(v interface{}) error{
	if v == nil{
		return errors.New("atomic value: nil value")
	}
	t := reflect.TypeOf(v)
	if t != av.t{
		return fmt.Errorf("atomic value: wrong type: %s", t)
	}
	av.v.Store(v)
	return nil
}

func (av *atomicValue) Load() interface{}{
	return av.v.Load()
}

func (av *atomicValue) TypeOfValue() reflect.Type{
	return av.t
}