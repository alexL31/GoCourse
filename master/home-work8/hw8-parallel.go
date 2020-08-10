// Предложение по применению -проверка работоспособности
//web-сервера, например. Или изменения содержимого...

package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	urls := []string{
		"http://www.geekbrains.ru",
		"http://www.yandex.ru",
		"http://www.hp.com",
		"http://youtube.com",
		"http://mail.ru",
		"http://google.com",
		"http://facebook.com",
		"http://aviasales.ru",
		"http://rbc.ru",
		"http://cian.ru"}
	start := time.Now()
	ch := make(chan string)
	for _, url := range urls { //os.Args[1:] {
		go fetch(url, ch)
	}
	for range urls { //os.Args[1:] {
		fmt.Print(<-ch) // receive from channel
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body) // to devNull
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	ch <- fmt.Sprintf("%.2fs %7d %s\n", time.Since(start).Seconds(), nbytes, url)
}
