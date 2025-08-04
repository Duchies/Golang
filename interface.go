package main

import (
	"fmt"
)

func main() {
	fmt.Println("Main started...")

	circleObj := circle{
		radius: 10,
	}

	squareObj := square{
		length: 5,
	}
	// without interface
	fmt.Println("Circle Area", circleObj.getArea())
	fmt.Println("square Area", squareObj.getArea())

	// with interface
	arr := []shape{&circleObj, &squareObj}

	for _, obj := range arr {
		fmt.Println("Shape Area:", obj.getArea())
	}

	fmt.Println("Main finished...")
}

type circle struct {
	radius int
}

type square struct {
	length int
}

func (c *circle) getArea() int {
	return 2 * c.radius
}

func (s *square) getArea() int {
	return s.length * s.length
}

type shape interface {
	getArea() int
}

```
FollowUp Question
```
🔍 How does the shape interface identify which function to call?
Let’s look at this part of your code:

type shape interface {
	getArea() int
}
This says: "Any type that has a method getArea() int is a shape."

Best way to learn and remember


2) interface composition 

// Base interface
type Reader interface {
    Read() string
}

// Another base interface
type Writer interface {
    Write(string)
}

// Composed interface that embeds both Reader and Writer
type ReadWriter interface {
    Reader
    Writer
}

i have this interface 
now assume there is fucntion which accept the ReadWriter interface type in this case how we can call that fucntion
To call a function that accepts a ReadWriter interface, you need to pass a concrete type that implements both Reader and Writer interfaces. Here's a complete example:

go
package main

import "fmt"

// Base interfaces
type Reader interface {
    Read() string
}

type Writer interface {
    Write(string)
}

// Composed interface
type ReadWriter interface {
    Reader
    Writer
}

// Concrete implementation
type File struct {
    content string
}

func (f *File) Read() string {
    return f.content
}

func (f *File) Write(s string) {
    f.content = s
}

// Function that accepts ReadWriter
func ProcessData(rw ReadWriter, data string) {
    fmt.Println("Initial content:", rw.Read())
    rw.Write(data)
    fmt.Println("Updated content:", rw.Read())
}

func main() {
    // Create a concrete instance
    myFile := &File{content: "initial data"}
    
    // Call the function with our File instance
    ProcessData(myFile, "new data written to file")
    
    // Alternatively, you can create and pass anonymously
    ProcessData(&File{content: "another file"}, "more data")
}
Key points:
The File type implements both Read() and Write() methods, so it satisfies the ReadWriter interface.

The ProcessData function can accept:

Any concrete type that implements ReadWriter

Any value that implements both Reader and Writer

Either a pointer or value receiver, as long as the method set is satisfied (in this case we use pointer receiver)

When calling the function, you can pass:

A pre-created instance (myFile)

An anonymously created instance (&File{...})

This demonstrates how interface composition allows you to write functions that work with types that implement multiple behaviors.

can interface can be type of var in golang?
Using Interfaces as Variable Types in Go
Yes, in Go you can declare variables with interface types. This is one of the fundamental uses of interfaces in Go - to specify that a variable can hold any value that implements the interface's methods.

Basic Example
go
package main

import "fmt"

type Speaker interface {
    Speak() string
}

type Dog struct{}
func (d Dog) Speak() string { return "Woof!" }

type Cat struct{}
func (c Cat) Speak() string { return "Meow!" }

func main() {
    // Declaring a variable with interface type
    var animal Speaker
    
    // Assigning different concrete types
    animal = Dog{}
    fmt.Println(animal.Speak()) // "Woof!"
    
    animal = Cat{}
    fmt.Println(animal.Speak()) // "Meow!"
}
Key Characteristics
Interface Variables Hold (value, type) Pairs:

When you assign a concrete value to an interface variable, Go stores both the value and its type

Nil Interfaces:

An interface variable is nil when both its value and type are nil

You can check this with if animal == nil

Empty Interface:

interface{} (or any in Go 1.18+) can hold any type

Often used for generic programming

Practical Example with Empty Interface
go
func printAnything(v interface{}) {
    fmt.Printf("Type: %T, Value: %v\n", v, v)
}

func main() {
    printAnything(42)         // int
    printAnything("hello")    // string
    printAnything(3.14)       // float64
    printAnything([]int{1,2}) // []int
}
Type Assertions
When you need to get the concrete value back:

go
func processAnimal(a Speaker) {
    // Type assertion
    if dog, ok := a.(Dog); ok {
        fmt.Println("It's a dog:", dog.Speak())
    } else if cat, ok := a.(Cat); ok {
        fmt.Println("It's a cat:", cat.Speak())
    }
}
Type Switches
For handling multiple possible types:

go
func describe(i interface{}) {
    switch v := i.(type) {
    case int:
        fmt.Printf("Integer: %v\n", v)
    case string:
        fmt.Printf("String: %v\n", v)
    default:
        fmt.Printf("Unknown type: %T\n", v)
    }
}
Interface variables are a powerful feature in Go that enable polymorphism and flexible code design.
