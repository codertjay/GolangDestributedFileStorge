package p2p

// HandshakeFunc is a function that is called when a new connection is established.
type HandshakeFunc func(any) error

func NOPHandshakeFunc(any) error {
	return nil
}
