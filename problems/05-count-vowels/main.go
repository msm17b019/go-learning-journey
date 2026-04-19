package main

import "fmt"

func main() {
	fmt.Println("Enter a string:")
	var word string
	fmt.Scan(&word)

	runes := []rune(word)

	counter := 0

	for _, ch := range runes {
		switch ch {
		case 'a', 'A', 'i', 'I', 'e', 'E', 'o', 'O', 'u', 'U':
			counter++
		}
	}

	fmt.Println("Total", counter, "vowels are there in the entered string.")
}
