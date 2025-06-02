package main

import "fmt"

func average(values ...float64) float64 {
	var result float64
	for _, value := range values {
		result += value
	}
	return result / float64(len(values))
}

func main() {
	fmt.Printf("Média dos valores %.2f\n", average(5, 6, 4))
}
