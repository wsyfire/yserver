package datapack

import (
	"yserver/message"
)

type Head interface {
	GetHeadSize() uint32
}

type Packet interface {
	Head
	Pack(msg message.Message) ([]byte, error)
	Unpack(data []byte) (message.Message, error)
}
