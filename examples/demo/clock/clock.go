package main

import (
	"io"
	"log"
	"net"
	"strconv"
	"time"
)

var sessionIdMap = make(map[int]net.Conn)
var sessionConnMap = make(map[net.Conn]int)
var sessionId int

func main() {

	lister, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal("listen err")
	}

	for {
		conn, err := lister.Accept()
		if err != nil {
			log.Print("accept err")
			continue
		}

		sessionId++
		log.Print("client", sessionId)
		sessionIdMap[sessionId] = conn
		sessionConnMap[conn] = sessionId
		go handleConn(conn, sessionId)
	}
}

func handleConn(conn net.Conn, sessionId int) {
	defer conn.Close()

	for {
		_, err := io.WriteString(conn, strconv.Itoa(sessionId)+"---"+time.Now().String()+"\n")
		if err != nil {
			return
		}

		time.Sleep(1 * time.Second)
	}
}
