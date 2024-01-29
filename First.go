package main

import (
	"fmt"
	"math"
)

// Радиус круга для области D
const radius = 5.0

// isInUpperRightQuarterCircle проверяет, принадлежит ли точка (x, y) верхней правой четверти круга радиуса 5 с центром в начале координат.
func isInUpperRightQuarterCircle(x, y float64) bool {
	return x >= 0 && y >= 0 && x*x+y*y <= radius*radius
}

// computeFunctionValue вычисляет значение функции для заданного x.
func computeFunctionValue(x float64) float64 {
	// Вычисление каждого из трех терминов функции.
	arcos := math.Pow(math.Acos(math.Pow(x, 3)/(math.Pow(x, 3)+1)), 3)
	underSqrt := math.Sqrt(math.Abs(x))
	secondTerm := math.Pow(underSqrt, 1.0/7.0)
	inSquare := math.Pow(5, math.Tan(x)) + math.Pow(math.Abs(x), math.Sin(x))
	thirdTerm := math.Pow(math.Log(inSquare)/math.Log(2), 2)

	// Суммирование терминов для получения общего значения функции.
	result := arcos + secondTerm + thirdTerm
	return result
}

func main() {
	var x float64

	// Ввод значения x с клавиатуры.
	fmt.Print("Введите значение x: ")
	fmt.Scan(&x)

	// Вычисление значения функции для введенного x.
	y := computeFunctionValue(x)

	// Проверка принадлежности точки области D.
	inDomain := isInUpperRightQuarterCircle(x, y)

	// Вывод результатов.
	fmt.Printf("Значение функции в точке (%.2f): %.4f\n", x, y)
	fmt.Printf("Принадлежит области D: %t\n", inDomain)
}
