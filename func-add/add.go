package main

import "fmt"

func add(x, y int) int {
	// add two int values and return the result
	return x + y
}

func sayHello() {
	// just print `hello`
	fmt.Println("hello")
}

func main() {
	// main function for adding logic
	var result int = add(42, 13)
	fmt.Println(result)
	sayHello()
}
