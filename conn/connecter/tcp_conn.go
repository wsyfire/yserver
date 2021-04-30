package connecter

import "net"

type TCPConn struct {
	addr   string
	conn   net.Conn
	connID uint32
}

func NewTCPConn(connID uint32) *TCPConn {
	return &TCPConn{connID: connID}
}

func (t *TCPConn) Dial(addr string) error {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return err
	}

	t.conn = conn
	t.addr = addr
	return nil
}

func (t *TCPConn) Stop() {
	t.conn.Close()
}

func (t *TCPConn) ConnID() uint32 {
	return t.connID
}

func (t *TCPConn) LocalAddr() string {
	if t.conn == nil {
		return ""
	}

	return t.conn.LocalAddr().String()
}

func (t *TCPConn) RemoteAddr() string {
	if t.conn == nil {
		return ""
	}

	return t.conn.RemoteAddr().String()
}

func (t *TCPConn) SetConn(conn net.Conn) {
	t.conn = conn
}

func (t *TCPConn) Write(data []byte) (int, error) {
	return t.conn.Write(data)
}

func (t *TCPConn) Read(data []byte) (n int, err error) {
	return t.conn.Read(data)
}
