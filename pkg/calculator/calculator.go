package calculator

import "time"

func Multiply(numbers []float64) float64 {
	time.Sleep(1 * time.Second) // Sleep only for new calculations
	result := 1.0
	for _, num := range numbers {
		result *= num
	}
	return result
}

func Plus(numbers []float64) float64 {
	time.Sleep(1 * time.Second) // Sleep only for new calculations
	result := 0.0
	for _, num := range numbers {
		result += num
	}
	return result
}
