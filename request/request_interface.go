package request

import (
	"yserver/conn/connecter"
	"yserver/message"
)

type Request interface {
	GetConn() connecter.Conn
	GetMessage() message.Message
}
