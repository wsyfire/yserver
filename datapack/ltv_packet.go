package datapack

import (
	"bytes"
	"encoding/binary"
	"yserver/message"
	"yserver/utils"
)

// ------------------------
// | msgID_len | data_len | real_data
// ------------------------
type LVTPacketHead struct {
	MsgIDSize uint32
	DataSize  uint32
}

type LTVPacket struct {
	LVTPacketHead
	bigEndian  bool
	maxDataLen uint32
}

func NewLTVPacket(msgIDSize, dataSize uint32) *LTVPacket {
	return &LTVPacket{LVTPacketHead: LVTPacketHead{MsgIDSize: msgIDSize, DataSize: dataSize},
		maxDataLen: 4096,
	}
}

func (l *LTVPacket) SetMaxDataLen(length uint32) {
	l.maxDataLen = length
}

func (l *LTVPacket) SetByteOrder(bigEndian bool) {
	l.bigEndian = bigEndian
}

func (l *LTVPacket) GetHeadSize() uint32 {
	return l.MsgIDSize + l.DataSize
}

func (l *LTVPacket) Pack(msg message.Message) ([]byte, error) {
	if msg.GetDataLen()+l.GetHeadSize() > l.maxDataLen {
		return nil, ErrPacketSizeExceed
	}

	buffer := make([]byte, l.GetHeadSize()+msg.GetDataLen())
	utils.GetEndian(l.bigEndian).PutUint32(buffer[:l.MsgIDSize], msg.GetMsgID())
	utils.GetEndian(l.bigEndian).PutUint32(buffer[l.MsgIDSize:l.GetHeadSize()], msg.GetDataLen())
	copy(buffer[l.GetHeadSize()+1:], msg.GetData())

	return buffer, nil
}

func (l *LTVPacket) Unpack(data []byte) (message.Message, error) {
	buffer := bytes.NewReader(data)

	var msgId uint32
	if err := binary.Read(buffer, utils.GetEndian(l.bigEndian), &msgId); err != nil {
		return nil, err
	}

	var dataLen uint32
	if err := binary.Read(buffer, utils.GetEndian(l.bigEndian), &dataLen); err != nil {
		return nil, err
	}

	msg := &message.DataMessage{}
	msg.SetMsgID(msgId)
	msg.SetDataLen(dataLen)

	if dataLen+l.GetHeadSize() > l.maxDataLen {
		return nil, ErrPacketSizeExceed
	}

	// message data need use msg.DataLen to read.
	return msg, nil
}
