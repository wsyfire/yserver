package client

import (
	"io"
	"io/ioutil"
	"yserver/conn/connecter"
	"yserver/datapack"
	"yserver/logger"
	"yserver/message"
	"yserver/serialize"
)

type YClient struct {
	addr       string
	conn       connecter.Conn
	serializer serialize.Serializer
	dataPack   datapack.Packet
}

func NewClient(conn connecter.Conn, addr string, serializer serialize.Serializer, dataPack datapack.Packet) Client {
	return &YClient{conn: conn, addr: addr, serializer: serializer, dataPack: dataPack}
}

func (c *YClient) Start() error {
	err := c.conn.Dial(c.addr)
	if err != nil {
		logger.Log.Errorf("client connect addr:%v err:%v", c.addr, err)
		return err
	}

	//
	return nil
}

func (c *YClient) Stop() {
	c.conn.Stop()
}

func (c *YClient) GetConnID() uint32 {
	return c.conn.ConnID()
}

func (c *YClient) Send(msgId uint32, v interface{}) (int, error) {
	data, err := c.serializer.Marshal(v)
	if err != nil {
		return 0, err
	}

	msg := &message.DataMessage{}
	msg.SetMsgID(msgId)
	msg.SetData(data)
	msg.SetDataLen(uint32(len(data)))

	packData, err := c.dataPack.Pack(msg)
	if err != nil {
		return 0, err
	}

	return c.conn.Write(packData)
}

func (c *YClient) GetNextMessage() (message.Message, error) {
	header, err := ioutil.ReadAll(io.LimitReader(c.conn, int64(c.dataPack.GetHeadSize())))
	if err != nil {
		return nil, err
	}

	msg, err := c.dataPack.Unpack(header)
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(io.LimitReader(c.conn, int64(msg.GetDataLen())))
	if err != nil {
		return nil, err
	}

	msg.SetData(data)
	return msg, nil
}
