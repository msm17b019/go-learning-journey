package main

import "fmt"

func main() {
	fmt.Println("Enter a number:")
	var n int
	var sum int = 0
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		sum += i
	}

	fmt.Println("The sum of first", n, "number is", sum)
}
