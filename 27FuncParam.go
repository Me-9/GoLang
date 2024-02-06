package main

import "fmt"

func main()  {
	sayHelloTo("Wisman", "Roy")
}

func sayHelloTo(firstName string, lastName string)  {
	fmt.Println("Hello", firstName, lastName)
}