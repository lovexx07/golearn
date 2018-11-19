package main

import "strings"

func main() {
	var builder1 strings.Builder

	builder1.Grow(1)

	f1 := func(b strings.Builder) {
		//b.Grow(1)  //拿到的是副本，不允许操作
	}
	f1(builder1)

	ch1 := make(chan strings.Builder, 1)
	ch1 <- builder1

	builder2 := <- ch1
	//builder2.Grow(1) //拿到的是副本，会报错
	_ = builder2


	builder3 := builder1
	//builder3.Grow(1) //副本不可调用
	_ = builder3


	f2 := func(bp *strings.Builder) {
		(*bp).Grow(1)  //这里不会引发panic, 但不是并发安全的
		builder4 := *bp
		//builder4.Grow(1)  //这里报错
		_ = builder4
	}
	f2(&builder1)

	builder1.Reset()
	builder5 := builder1
	builder5.Grow(1)  //未使用过的副本可以调用该方法
}