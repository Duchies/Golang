package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Main started...")
	ch := make(chan bool)
	for i := 1; i <= 5; i++ {
		go worker(i, ch)
		//<-ch
	}

	fmt.Println("Main finished...")
}

func worker(id int, check chan bool) {
	fmt.Printf("Worker %d started...\n", id)

	// work time
	time.Sleep(time.Second)

	fmt.Printf("Worker %d finished...\n", id)
	check <- true
}
