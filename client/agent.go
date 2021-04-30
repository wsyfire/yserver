package client

import (
	"io"
	"io/ioutil"
	"yserver/conn/connecter"
	"yserver/datapack"
	"yserver/message"
	"yserver/serialize"
)

type Agent struct {
	serializer serialize.Serializer
	dataPack   datapack.Packet
	conn       connecter.Conn
}

func NewAgent(conn connecter.Conn, serializer serialize.Serializer, dataPack datapack.Packet) Client {
	return &Agent{conn: conn, serializer: serializer, dataPack: dataPack}
}

func (a *Agent) Start() error {
	return nil
}

func (a *Agent) Stop() {
	a.conn.Stop()
}

func (a *Agent) GetConnID() uint32 {
	return a.conn.ConnID()
}

func (a *Agent) Send(msgId uint32, v interface{}) (int, error) {
	data, err := a.serializer.Marshal(v)
	if err != nil {
		return 0, err
	}

	msg := &message.DataMessage{}
	msg.SetMsgID(msgId)
	msg.SetData(data)
	msg.SetDataLen(uint32(len(data)))

	packData, err := a.dataPack.Pack(msg)
	if err != nil {
		return 0, err
	}

	return a.conn.Write(packData)
}

func (a *Agent) GetNextMessage() (message.Message, error) {
	header, err := ioutil.ReadAll(io.LimitReader(a.conn, int64(a.dataPack.GetHeadSize())))
	if err != nil {
		return nil, err
	}

	msg, err := a.dataPack.Unpack(header)
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(io.LimitReader(a.conn, int64(msg.GetDataLen())))
	if err != nil {
		return nil, err
	}

	msg.SetData(data)
	return msg, nil
}
