package p2p

import (
	"fmt"
	"net"
	"sync"
)

// TCPPeer represents the remote node over a TCP established connection.
type TCPPeer struct {
	conn net.Conn

	// if we dail and a connection => outbound == true
	// if we accept and retrieve connection => outbound == false
	outbound bool
}

type TCPTransport struct {
	ListenAddress string
	Listener      net.Listener
	mu            sync.Mutex
	peers         map[net.Addr]Peer
}

func NewTCPTransport(listenAddress string) *TCPTransport {
	return &TCPTransport{
		ListenAddress: listenAddress,
		peers:         make(map[net.Addr]Peer),
	}
}

func (t *TCPTransport) ListenAndAccept() error {
	var err error
	t.Listener, err = net.Listen("tcp", t.ListenAddress)
	if err != nil {
		return err
	}
	go t.startAcceptLoop()

	return nil
}

func (t *TCPTransport) startAcceptLoop() {
	for {
		conn, err := t.Listener.Accept()
		if err != nil {
			fmt.Printf("TCP accpet error: %s\n", err)
		}

		go t.handleConn(conn)

	}

}

func (t *TCPTransport) handleConn(conn net.Conn) {
	fmt.Printf("New connection from %s\n", conn.RemoteAddr())
}
