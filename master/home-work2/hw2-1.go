package main

import (
	"fmt"
)

func main() {
	n := new(int)
	fmt.Println("Введите целое число")
	fmt.Scanln(n)
	if checkOdd(*n) {
		fmt.Println("Четное")
	} else {
		fmt.Println("Нечетное")
	}

	if checkDevideThree(*n) {
		fmt.Println("Делится на 3")
	} else {
		fmt.Println("Не делится на 3 без остатка")
	}

	//firstNum := enterNumber("Введите первое число: ")

}
func checkOdd(number int) bool {
	if number%2 == 0 {
		return true
	} else {
		return false
	}
}
func checkDevideThree(number int) bool {
	if number%3 == 0 {
		return true
	} else {
		return false
	}
}
