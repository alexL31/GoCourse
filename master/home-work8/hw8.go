// Исследуйте работу последовательного и параллельного
// сканера веб-сайтов, задав большое (не менее 10)
// количество URL. Какие выводы можно сделать?

package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
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
	for _, url := range urls { //os.Args[1:] {
		fetch(url)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	//bytes, err := ioutil.ReadAll(resp.Body)
	nbytes, err := io.Copy(ioutil.Discard, resp.Body) // to devNull
	resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%.2fs %7d %s\n", time.Since(start).Seconds(), nbytes, url)
}
