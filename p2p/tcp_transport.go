package p2p

import (
	"net"
	"sync"
)

type TCPTransport struct {
	listenAddress string
	listener      net.Listener

	// mutex above the thing u want to protect
	mu    sync.RWMutex
	peers map[net.Addr]Peer
}

func NewTCPTransport(listenAddr string) Transport {
	return &TCPTransport{
		listenAddress: listenAddr,
		//peers:         make(map[net.Addr]Peer),
	}
}

func Tset() {
	t := NewTCPTransport(":4344").(*TCPTransport)

	t.listener.Accept()
}
