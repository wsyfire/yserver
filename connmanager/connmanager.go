package connmanager

import (
	"sync"
	"yserver/conn/connecter"
	"yserver/logger"
)

type ConnManager struct {
	connMap  map[uint32]connecter.Conn
	connLock sync.RWMutex
}

func NewConnManager() *ConnManager {
	return &ConnManager{connMap: make(map[uint32]connecter.Conn)}
}

func (c *ConnManager) Add(connID uint32, conn connecter.Conn) {
	c.connLock.Lock()
	defer c.connLock.Unlock()

	if _, ok := c.connMap[connID]; ok {
		logger.Log.Warnf("id:%d is exist in conn map.\n", connID)
	}

	c.connMap[connID] = conn
}

func (c *ConnManager) Get(connID uint32) connecter.Conn {
	c.connLock.RLock()
	defer c.connLock.RUnlock()

	if v, ok := c.connMap[connID]; ok {
		return v
	}

	return nil
}

func (c *ConnManager) Remove(connID uint32) {
	c.connLock.Lock()
	defer c.connLock.Unlock()

	if _, ok := c.connMap[connID]; ok {
		delete(c.connMap, connID)
	}
}

func (c *ConnManager) RemoveConn(conn connecter.Conn) {
	c.connLock.Lock()
	defer c.connLock.Unlock()

	if _, ok := c.connMap[conn.ConnID()]; ok {
		delete(c.connMap, conn.ConnID())
	}
}

func (c *ConnManager) ClearConn(conn connecter.Conn) {
	c.connLock.Lock()
	defer c.connLock.Unlock()

	if _, ok := c.connMap[conn.ConnID()]; ok {
		conn.Stop()

		delete(c.connMap, conn.ConnID())
	}
}

func (c *ConnManager) ClearAllConn() {
	c.connLock.Lock()
	defer c.connLock.Unlock()

	for connID, conn := range c.connMap {
		conn.Stop()

		delete(c.connMap, connID)
	}
}

func (c *ConnManager) Len() int {
	return len(c.connMap)
}
