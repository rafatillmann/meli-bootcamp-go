package main

import (
	"errors"
	"fmt"
	"math"
)

const (
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
)

func minOperation(values ...float64) float64 {
	result := values[0]
	for _, value := range values[1:] {
		result = math.Min(result, value)
	}
	return result
}

func averageOperation(values ...float64) float64 {
	var result float64
	for _, value := range values {
		result += value
	}
	return result / float64(len(values))
}

func maxOperation(values ...float64) float64 {
	result := values[0]
	for _, value := range values[1:] {
		result = math.Max(result, value)
	}
	return result
}

func operation(op string) (func(values ...float64) float64, error) {
	switch op {
	case minimum:
		return minOperation, nil
	case average:
		return averageOperation, nil
	case maximum:
		return maxOperation, nil
	default:
		return nil, errors.New("Operation invalid")
	}
}

func main() {

	minFunc, err := operation(minimum)
	if err != nil {
		fmt.Println(err.Error())
        return
	}
	averageFunc, err := operation(average)
	if err != nil {
		fmt.Println(err.Error())
        return
	}
	maxFunc, err := operation(maximum)
	if err != nil {
		fmt.Println(err.Error())
        return
	}
	minValue := minFunc(2, 3, 3, 4, 10, 2, 4, 5)
	averageValue := averageFunc(2, 3, 3, 4, 1, 2, 4, 5)
	maxValue := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)

	fmt.Printf("Valor mínimo %.2f\n", minValue)
	fmt.Printf("Média %.2f\n", averageValue)
	fmt.Printf("Valor máximo %.2f\n", maxValue)
}
