package listener

import (
	"net"
	"yserver/conn/connecter"
)

type TCPListener struct {
	addr      string
	listener  net.Listener
	running   bool
	connChan  chan connecter.Conn
	curConnID uint32
}

func NewTCPListener(addr string) *TCPListener {
	return &TCPListener{addr: addr,
		running:  false,
		connChan: make(chan connecter.Conn)}
}

func (t *TCPListener) ListenAndServe() {
	l, err := net.Listen("tcp", t.addr)
	if err != nil {
		panic(err)
	}

	t.listener = l
	t.running = true

	t.serve()
}

func (t *TCPListener) GetAddr() string {
	return t.addr
}

func (t *TCPListener) serve() {
	defer t.Stop()

	for t.running {
		conn, err := t.listener.Accept()
		if err != nil {
			continue
		}

		t.curConnID++
		newConn := connecter.NewTCPConn(t.curConnID)
		newConn.SetConn(conn)

		t.connChan <- newConn
	}
}

func (t *TCPListener) Stop() {
	if t.running {
		t.running = false
		t.listener.Close()
		close(t.connChan)
	}
}

func (t *TCPListener) GetNewConn() chan connecter.Conn {
	return t.connChan
}
