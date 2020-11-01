package receiver

import (
	"github.com/mdlayher/ethernet"
	"log"
	"net"
)

// Frame size - Maximum Transition Unit
const MTU int = 800

// Receives and processes messages sent from other nodes
func ReceiveAndProcess(c net.PacketConn) {
	var f ethernet.Frame
	b := make([]byte, MTU)

	// Keep receiving messages forever.
	for {
		n, addr, err := c.ReadFrom(b)
		if err != nil {
			log.Fatalf("failed to receive message: %v", err)
		}

		// Process stage
		// Unpack Ethernet II frame into Go representation.
		if err := (&f).UnmarshalBinary(b[:n]); err != nil {
			log.Fatalf("failed to unmarshal protocol frame: %v", err)
		}

		// TODO save to temp files
		log.Printf("[%s] %s", addr.String(), string(f.Payload))
	}
}
