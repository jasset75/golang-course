package main

import "fmt"

func add(x, y int) int {
	// add two int numbers and return the result
	return x + y
}

func sayHello() {
	// just print `hello`
	fmt.Println("hello")
}

func main() {
	fmt.Println(add(42, 13))
	sayHello()
}
