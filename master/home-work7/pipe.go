// Перепишите программу-конвейер, ограничив количество
// передаваемых для обработки значений и обеспечив корректное
// завершение всех горутин.
package main

import (
	"fmt"
)

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// генерация
	go func() {
		for x := 0; ; x++ {
			naturals <- x
		}
	}()

	// возведение в квадрат
	go func() {
		for {
			x := <-naturals
			squares <- x * x
		}
	}()

	// печать
	for {
		fmt.Println(<-squares)
	}
}
