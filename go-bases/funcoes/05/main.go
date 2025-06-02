package main

import (
	"errors"
	"fmt"
)

const (
	Dog       = "dog"
	Cat       = "cat"
	Hamster   = "hamster"
	Tarantula = "tarantula"
)

func animalDog(value int) float64 {
	return 10 * float64(value)
}

func animalCat(value int) float64 {
	return 5 * float64(value)
}

func animalHamster(value int) float64 {
	return 0.25 * float64(value)
}

func animalTarantula(value int) float64 {
	return 0.15 * float64(value)
}

func animal(animal string) (func(value int) float64, error) {
	switch animal {
	case Dog:
		return animalDog, nil
	case Cat:
		return animalCat, nil
	case Hamster:
		return animalHamster, nil
	case Tarantula:
		return animalTarantula, nil
	default:
		return nil, errors.New("Unexpected error")
	}
}

func main() {

	animalDog, err := animal(Dog)
	if err != nil {
		fmt.Println(err)
	}
	animalCat, err := animal(Cat)
	if err != nil {
		fmt.Println(err)
	}
	animalHamster, err := animal(Hamster)
	if err != nil {
		fmt.Println(err)
	}
	animalTarantula, err := animal(Tarantula)
	if err != nil {
		fmt.Println(err)
	}

	result := animalDog(2) + animalCat(3) + animalHamster(1) + animalTarantula(1)
	fmt.Printf("Quantia de alimento é %.3f kg\n", result)

}
