package main

import (
	"fmt"
	"time"
	"yserver/client"
	"yserver/conn/connecter"
	"yserver/datapack"
	"yserver/serialize/json"
)

type Hello struct {
	Name string `json:"name"`
}

func main() {
	for i := uint32(0); i < 1; i++ {
		conn := connecter.NewTCPConn(i)
		c := client.NewClient(conn, ":8787", json.NewSerializer(), datapack.NewLTVPacket(4, 4))
		err := c.Start()
		if err != nil {
			fmt.Println(err)
		}

		c.Send(1, Hello{Name: "YYY"})
		//time.Sleep(500 * time.Millisecond)
	}
	time.Sleep(5 * time.Second)
}
