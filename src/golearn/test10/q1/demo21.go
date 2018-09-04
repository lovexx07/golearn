package main

func main() {
	ch1 := make(chan int, 1)
	ch1 <-1

//	ch1 <-2  //通道已满，阻塞

	ch2 := make(chan int, 1)
	//elem, ok := <-ch2//通道以及空了 阻塞
	//_, _ = elem, ok
	ch2 <-1

	var ch3 chan int //定义通道
	//ch3 <- 1 //通道为nil , 因此会造成永久阻塞
	//<-ch3

	_ = ch3

}