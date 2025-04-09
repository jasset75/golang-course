package main

import (
	"complex/pkg/utils"
	"fmt"
)

func main() {
	fmt.Println("complex rotation...")

	var zNumber complex128 = complex(2, 1)
	fmt.Printf("z = %.2f\n", zNumber)

	var zRotated complex128 = utils.RotateDegrees(zNumber, 45.0)
	fmt.Printf("z rotated = %.2f\n", zRotated)
}
