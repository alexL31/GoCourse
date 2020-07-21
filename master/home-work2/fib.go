package main

import (
	"fmt"
	"math/big"
)

func main() {

	y := big.NewInt(0)
	fmt.Printf("%v\n", y)
	x := big.NewInt(1)
	for i := 0; i < 100; i++ {

		x.Add(x, y)
		y.Sub(x, y)

		fmt.Printf("%v\n", y)
	}
}
