package main

import (
	"log"
	"math"
)

func main() {
	log.Println(eratosthenes(10000))
}

func eratosthenes(n int) []int {
	arr := make([]bool, n)
	count := 0
	// var arr []bool
	for i := 2; i <= int(math.Sqrt(float64(n)+1)); i++ {
		if arr[i] == false {
			for j := i * i; j < n; j += i {
				arr[j] = true
				count++
			}
		}
	}
	var primes []int

	for i, isComposite := range arr {
		if i > 1 && !isComposite {
			primes = append(primes, i)
		}
	}

	return primes
}
