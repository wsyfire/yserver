package client

import "yserver/message"

type Client interface {
	Start() error
	Stop()
	Send(msgId uint32, v interface{}) (int, error)
	GetConnID() uint32
	GetNextMessage() (message.Message, error)
}
