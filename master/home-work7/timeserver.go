// Добавьте для time-сервера возможность его завершения
// при вводе команды exit.
package main

import (
	"io"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	cancel := make(chan int)
	go func() {
		os.Stdin.Read(make([]byte, 1)) // read a single byte
		cancel <- 1
	}()

	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n\r"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
