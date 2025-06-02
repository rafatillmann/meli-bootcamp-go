package main

import "fmt"

func main() {
	var word string
	fmt.Print("Digite alguma palavra: ")
	fmt.Scan(&word)

	fmt.Println(len(word))

	for _, letter := range word {
		fmt.Printf("Letter: %c, Value: %d\n", letter, letter)
	}
}
