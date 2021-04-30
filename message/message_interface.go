package message

type Message interface {
	GetData() []byte
	GetDataLen() uint32
	GetMsgID() uint32

	SetData(data []byte)
	SetDataLen(len uint32)
	SetMsgID(msgID uint32)
}
