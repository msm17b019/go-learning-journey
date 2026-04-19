package main

import (
	"fmt"
)

func main() {
	fmt.Println("Enter a number:")
	var n int
	fmt.Scan(&n)

	a, b := 0, 1

	for i := 0; i < n; i++ {
		fmt.Print(a, " ")
		next := a + b
		b = a
		a = next
	}
}
