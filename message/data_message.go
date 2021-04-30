package message

type DataMessage struct {
	msgID   uint32
	data    []byte
	dataLen uint32
}

func (m *DataMessage) GetData() []byte {
	return m.data
}

func (m *DataMessage) GetDataLen() uint32 {
	return m.dataLen
}

func (m *DataMessage) GetMsgID() uint32 {
	return m.msgID
}

func (m *DataMessage) SetData(data []byte) {
	m.data = data
}

func (m *DataMessage) SetDataLen(len uint32) {
	m.dataLen = len
}

func (m *DataMessage) SetMsgID(msgID uint32) {
	m.msgID = msgID
}
