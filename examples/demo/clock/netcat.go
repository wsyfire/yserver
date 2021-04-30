package main

import (
	"io"
	"log"
	"net"
	"os"
	"strconv"
	"sync"
)

func main() {
	count := 1
	if len(os.Args) > 1 {
		count, _ = strconv.Atoi(os.Args[1])
	}

	var wg sync.WaitGroup

	for i := 0; i < count; i++ {
		wg.Add(1)
		go handleClient(wg)
	}

	wg.Wait()
}

func handleClient(wg sync.WaitGroup) {
	defer wg.Done()

	client, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
		return
	}

	defer client.Close()

	if _, err := io.Copy(os.Stdout, client); err != nil {
		log.Fatal(err)
	}
}
