package main

import (
	"fmt"
)

func swap(strA, strB string) (string, string) {
	return strB, strA
}

func main() {
	var strA, strB string

	fmt.Println("Swap two strings...")

	fmt.Printf("\t enter Str A: ")
	fmt.Scanln(&strA)

	fmt.Printf("\t enter Str B: ")
	fmt.Scanln(&strB)

	strA, strB = swap(strA, strB)

	fmt.Printf("Str A: %s, Str B: %s\n", strA, strB)

}
