package connecter

type Conn interface {
	Dial(addr string) error
	Stop()
	ConnID() uint32

	LocalAddr() string
	RemoteAddr() string

	Write(data []byte) (int, error)
	Read(data []byte) (n int, err error)
}
