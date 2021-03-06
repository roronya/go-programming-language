package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

var port = flag.Int("port", 8000, "port")

func main() {
	flag.Parse()
	address := fmt.Sprintf("localhost:%d", *port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%s is listening...", address)
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // 例: 接続が切れた
			continue
		}
		go handleConn(conn) // 接続を並行して処理する
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return // 例: クライアントとの接続が切れた
		}
		time.Sleep(1 * time.Second)
	}
}
