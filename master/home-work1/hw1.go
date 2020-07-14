package main

import (
	"fmt"
	"math"
)

const a = 3
const b = 5

var c float64
var summa = new(float64)

func main() {
	c = math.Sqrt(a*a + b*b)
	fmt.Println("Гипотенуза треугольника = ", c)
	fmt.Println("Площадь треугольника = ", a*b)
	fmt.Println("Периметр треугольника = ", a+b+c)

	fmt.Println("Введите сумму вклада")
	fmt.Scanln(&summa)
}
