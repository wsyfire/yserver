package server

import (
	"os"
	"os/signal"
	"syscall"
	"yserver/client"
	"yserver/conn/connecter"
	"yserver/conn/listener"
	"yserver/connmanager"
	"yserver/datapack"
	"yserver/logger"
	"yserver/msghandler"
	"yserver/serialize"
)

type Server struct {
	listener   listener.Listener
	serializer serialize.Serializer
	dataPack   datapack.Packet
	connMgr    *connmanager.ConnManager
	msgHandler *msghandler.MsgHandler
	stopChan   chan int
}

type Option func(s *Server)

func NewServer(listener listener.Listener, serializer serialize.Serializer,
	dataPack datapack.Packet, options ...Option) *Server {
	s := &Server{listener: listener,
		serializer: serializer,
		dataPack:   dataPack,
		connMgr:    connmanager.NewConnManager(),
		msgHandler: msghandler.NewMsgHandler(10),
		stopChan:   make(chan int)}

	for _, op := range options {
		op(s)
	}

	return s
}

func (s *Server) Run() {
	logger.Log.Infof("server %v is running...", s.listener.GetAddr())

	go func() {
		for newConn := range s.listener.GetNewConn() {
			logger.Log.Debugf("server got a new conn, connID:%d, ip:%v, conn:%p.",
				newConn.ConnID(), newConn.RemoteAddr(), newConn)

			go s.HandleConn(newConn)
		}
	}()

	go func() {
		s.listener.ListenAndServe()
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGKILL, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	select {
	case sig := <-ch:
		logger.Log.Infof("server got signal %v!", sig)
		close(s.stopChan)
	case <-s.stopChan:
		logger.Log.Warnln("server shut down!")
	}

	logger.Log.Warnln("server is stopping!")
	s.listener.Stop()
	s.connMgr.ClearAllConn()
}

func (s *Server) HandleConn(conn connecter.Conn) {
	s.connMgr.Add(conn.ConnID(), conn)

	agent := client.NewAgent(conn, s.serializer, s.dataPack)

	defer func() {
		s.connMgr.ClearConn(conn)

		logger.Log.Debugf("connID:%d(ip:%v) is disconnect, connNum:%v.",
			conn.ConnID(), conn.RemoteAddr(), s.connMgr.Len())
	}()

	for {
		msg, err := agent.GetNextMessage()
		if err != nil {
			//logger.Log.Debug("agent get next message err:", err)
			return
		}

		s.msgHandler.SendMessage(agent, msg)
	}
}

func (s *Server) Stop() {
	close(s.stopChan)
}
