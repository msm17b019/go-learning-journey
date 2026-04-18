package main

import "fmt"

func main() {
	var N int
	fmt.Println("Enter a number:")
	fmt.Scan(&N)

	if N%2 == 0 {
		fmt.Printf("%v is an even number\n", N)
	} else {
		fmt.Printf("%v is an odd number\n", N)
	}
}
