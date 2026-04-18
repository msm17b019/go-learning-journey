package main

import (
	"fmt"
)

func main() {
	fmt.Println("Enter a string:")
	var word string
	fmt.Scan(&word)

	runes := []rune(word)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	fmt.Println("The reversed string of", word, "is", string(runes))
}
