package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	fmt.Println("Planning a school trip")
	var wg sync.WaitGroup

	wg.Add(3)
	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println("Rahul has done its task")
		wg.Done()
	}()

	go func() {
		time.Sleep(3 * time.Second)
		fmt.Println("tafe has done its task")
		wg.Done()
	}()

	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("ziyo has done its task")
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("all task done")
}
