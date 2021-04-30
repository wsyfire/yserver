package main

import (
	"github.com/sirupsen/logrus"
	"yserver/conn/listener"
	"yserver/datapack"
	"yserver/logger"
	"yserver/serialize/json"
	"yserver/server"
)

func main() {
	l := logrus.New()
	l.SetLevel(logrus.DebugLevel)
	logger.SetLogger(l)

	app := server.NewServer(listener.NewTCPListener("localhost:8787"),
		json.NewSerializer(), datapack.NewLTVPacket(4, 4),
	)
	app.Run()
}
