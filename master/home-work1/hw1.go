package main

import (
	"fmt"
	"math"
)

const a = 3
const b = 5
const kurs = 71

var c float64
var summa = new(float64)

func main() {
	fmt.Println("Введите сумму в рублях для конвертации в доллары")
	fmt.Scanln(summa)
	fmt.Println("Вы можете приобрести", *summa/kurs)

	c = math.Sqrt(a*a + b*b)
	fmt.Println("Гипотенуза треугольника = ", c)
	fmt.Println("Площадь треугольника = ", a*b/2)
	fmt.Println("Периметр треугольника = ", a+b+c)

	fmt.Println("Введите сумму вклада в рублях")
	fmt.Scanln(summa)
	fmt.Println(*summa)
}
