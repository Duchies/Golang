package main

import (
	"fmt"
	"sync"
)

func main() {

	// channel value print
	// log : sent 1 to channel

	ch := make(chan int)

	wg := &sync.WaitGroup{}
	wg.Add(2)
	go func(ch chan int) {
		//time.Sleep(10 * time.Second)
		fmt.Println(<-ch)
		wg.Done()
	}(ch)

	go func(ch chan int) {
		ch <- 1
		fmt.Println("Sent 1 to channel")
		close(ch)
		wg.Done()
	}(ch)

	wg.Wait()
	fmt.Println("All goroutines finished")
	fmt.Println("Exiting main function")
}
