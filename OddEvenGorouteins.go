package main

import "fmt"

// program to find the odd even using go-routeins
func main() {

	var intSlice = []int{91, 42, 23, 14, 15, 76, 87, 28, 19, 95}
	chOdd := make(chan int)
	chEven := make(chan int)

	go odd(chOdd)
	go even(chEven)

	for _, val := range intSlice {
		if val%2 == 0 {
			chEven <- val
		} else {
			chOdd <- val
		}
	}

}

func odd(chOdd <-chan int) {
	for value := range chOdd {
		fmt.Println("this is odd value:", value)
	}
}

func even(chEven <-chan int) {
	for value := range chEven {
		fmt.Println("this is even value:", value)
	}
}
