package main

import "fmt"

func main() {
	word := "word"

	fmt.Println(len(word))

	for _, letter := range word {
		fmt.Printf("Letter: %c, Value: %d\n", letter, letter)
	}
}
