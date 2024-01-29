package main

import (
	"fmt"
	"math"
)

func computeFunctionValue(x float64) float64 {
	var result float64

	if x <= -1 {
		result = math.Pow(x, 6) * math.Log10(math.Abs(math.Pow(6, x)-math.Pow(math.Abs(x-5), x)))
	} else if x >= -1 && x < 1 {
		result = math.Sin(x/(1-2*x*x)) - 1
	} else if x >= 1 {
		result = math.Asin(1 / math.Pow(x, 4))
	}

	return result
}

func main() {
	var x float64
	fmt.Print("Введите значение x: ")
	fmt.Scan(&x)
	y := computeFunctionValue(x)
	fmt.Printf("Значение функции y(x) в точке %.2f: %.4f\n", x, y)
}
