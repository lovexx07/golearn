package main

import (
	"sync/atomic"
	"fmt"
	"time"
)

func main() {
	var count uint32
	trigger := func(i uint32, fn func()) {
		for {
			if n := atomic.LoadUint32(&count); n == i{
				fn()
				atomic.AddUint32(&count, 1) //原子操作，加以保护，避免多线程操作 竞态条件
				break
			}
			time.Sleep(time.Nanosecond) //休眠一下等待执行
		}
	}

	for i := uint32(0); i < 10; i++{
		go func(i uint32) {
			fn := func() {
				fmt.Println(i)
			}
			trigger(i, fn)
		}(i)
	}
	trigger(10, func() {

	})// 是主goruntione 最后一个运行完毕，最好传10，>10会进入循环， <10可能会发生未知终止
}