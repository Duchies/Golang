package main

import "fmt"

func main() {
	var name, age string

	// takes input value for name and age
	fmt.Print("Enter your name: ")
	fmt.Scan(&name)
	fmt.Println("Enter your age: ")
	fmt.Scan(&age)
	fmt.Printf("Hello %s you age is : %s!", name, age)

}
