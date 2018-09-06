package main

import "fmt"

var channels = [3]chan int{
	nil,
	make(chan int),//无缓冲通道
	nil,
}

var channls2 = [2]chan int{
	make(chan int, 1),
	make(chan int, 1),
}

var numbers = []int{1, 2, 3}

func main() {
	select {
	case getChan(0) <- getNumber(0):
		fmt.Println("The first candidate case is selected.")
	case getChan(1) <- getNumber(1):
		fmt.Println("The second condidate case is selected")
	case getChan(2) <- getNumber(2):
		fmt.Println("The third condidate case is selected")
	default:
		fmt.Println("No ocndidate case seledted")
	}
	fmt.Println("------------------")
	for i := 0; i < len(numbers); i++ {
		select {
		case getChan(0) <- getNumber(0):
			fmt.Println("The first candidate case is selected.")
		case getChan(1) <- getNumber(1):
			fmt.Println("The second condidate case is selected")
		default:
			fmt.Println("No ocndidate case seledted")
			break 
		}
		fmt.Printf("for loop %d", i)
	}

}
func getChan(i int) chan int {
	fmt.Printf("channels[%d]\n", i)
	return channels[i]
}

func getNumber(i int) int {
	fmt.Printf("numbers[%d]\n", i)
	return numbers[i]
}

func getChan2(i int) chan int {
	fmt.Printf("channels2[%d]\n", i)
	return channls2[i]
}

