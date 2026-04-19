package main

import (
	"fmt"
)

func main() {
	fmt.Println("Enter a number:")
	var n int
	fmt.Scan(&n)

	fact := 1

	if n < 0 {
		fmt.Println("Factorial is not defined for negative numbers.")
		return
	}

	for i := 1; i <= n; i++ {
		fact *= i
	}

	fmt.Printf("The factorial of %v is %v.\n", n, fact)
}
