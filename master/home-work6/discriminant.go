package main

import (
	"fmt"
	"math"
)

func discr(a, b, c float64) float64 {
	return b*b - 4*a*c
}

func main() {
	var (
		a float64 = 2
		b float64 = 100
		c float64 = 1
	)
	discr := discr(a, b, c)
	fmt.Println(discr)
	if discr > 0 {
		x1 := (-b + math.Sqrt(discr)) / (2 * a)
		x2 := (-b - math.Sqrt(discr)) / (2 * a)
		fmt.Println(x1, x2)
	} else if discr == 0 {
		x := -b / (2 * a)
		fmt.Println(x)
	} else {
		fmt.Println("Корней нет")
	}
}
