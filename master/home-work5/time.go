package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println(t.Format("02-01-2006"))
	fmt.Println(t.Format("02-01-2006 15:04:05"))
}

//Считаю полезным добавить в методичку моменты,изложенные в
//статье  Time in Go: A primer , связанные с локацией и
//часовыми поясами
