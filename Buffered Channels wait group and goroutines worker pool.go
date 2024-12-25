package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	fmt.Println("Buffered Channels wait group and goroutines making use as worker pool")

	tasks := []Task{
		{ID: 1, Name: "Task 1"},
		{ID: 2, Name: "Task 2"},
		{ID: 3, Name: "Task 3"},
		{ID: 4, Name: "Task 4"},
		{ID: 5, Name: "Task 5"},
		{ID: 6, Name: "Task 6"},
		{ID: 7, Name: "Task 7"},
		{ID: 8, Name: "Task 8"},
		{ID: 9, Name: "Task 9"},
		{ID: 10, Name: "Task 10"},
		{ID: 11, Name: "Task 11"},
		{ID: 12, Name: "Task 12"},
	}

	noOfWorker := 3

	taskCh := make(chan Task, len(tasks))
	resultCh := make(chan Result, len(tasks))

	var wg sync.WaitGroup

	for i := 1; i <= noOfWorker; i++ {
		wg.Add(1)
		go doWorkMachine(i, taskCh, resultCh, &wg)
	}

	for _, task := range tasks {
		taskCh <- task
	}
	close(taskCh)

	go func() {
		wg.Wait()
	}()

	for result := range resultCh {
		fmt.Printf("Task Completed manually verification : %s\n", result.outcome)
	}

	fmt.Println("task done")
}

func doWorkMachine(id int, tasks <-chan Task, results chan<- Result, wg *sync.WaitGroup) {
	// chan // read-write
	// <-chan // read-only
	// chan<- // write-only
	defer wg.Done()

	for task := range tasks {
		fmt.Printf("Worker started the worker ID: %d and its task is : %s\n", id, task.Name)
		// time to do work
		time.Sleep(time.Duration(1 * time.Second))
		output := fmt.Sprintf("Task done by worked ID: %d and its task is : '%s'", id, task.Name)
		results <- Result{task.ID, output}
	}

	close(results)

}

type Task struct {
	ID   int
	Name string
}

type Result struct {
	TaskID  int
	outcome string
}
