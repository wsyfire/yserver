package listener

import (
	"yserver/conn/connecter"
)

type Listener interface {
	ListenAndServe()
	GetAddr() string
	Stop()
	GetNewConn() chan connecter.Conn
}
