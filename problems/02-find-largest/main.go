package main

import "fmt"

func main() {
	fmt.Println("Enter 3 Numbers:")
	var x, y, z int
	fmt.Scan(&x, &y, &z)
	large := max(x, y, z)

	fmt.Println(large, "is the largest number.")

}
