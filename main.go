package main

import (
	"GolangDestributedFileStorge/p2p"
	"log"
)

func main() {
	tr := p2p.NewTCPTransport(":4000")
	err := tr.ListenAndAccept()
	if err != nil {
		log.Fatal(err)
	}

	select {}

}
