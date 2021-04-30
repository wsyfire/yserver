package msghandler

import (
	"yserver/client"
	"yserver/logger"
	"yserver/message"
)

type MsgCallback func(client.Client, message.Message)

type HandlerValue struct {
	Val interface{}
	Cb  MsgCallback
}

type MsgHandler struct {
	handlerMap     map[uint32]*HandlerValue
	workerPoolSize uint32
	workerPool     []chan MsgCallback
}

func NewMsgHandler(workerPoolSize uint32) *MsgHandler {
	return &MsgHandler{handlerMap: make(map[uint32]*HandlerValue),
		workerPoolSize: workerPoolSize}
}

func (m *MsgHandler) Register(msgID uint32, v interface{}, cb MsgCallback) {
	if _, ok := m.handlerMap[msgID]; ok {
		logger.Log.Warnln("don't duplication register msgID:%d", msgID)
	}

	m.handlerMap[msgID] = &HandlerValue{Val: v, Cb: cb}
}

func (m *MsgHandler) SendMessage(client client.Client, msg message.Message) {
	//workerId := client.GetConnID() % m.workerPoolSize
	//
	//m.workerPool[workerId] <-
}
