package ethernet

import (
	"log"
	"net"
	"time"

	"github.com/mdlayher/ethernet"
)

const etherType ethernet.EtherType = 0xcccc
const broadcast net.HardwareAddr = []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff}

type addr net.HardwareAddr

func send(conn net.PacketConn, dest net.HardwareAddr, source net.HardwareAddr, msg string) {
	// Message is broadcast to all machines in same network segment.
	f := &ethernet.Frame{
		Destination: ethernet.Broadcast,
		Source:      source,
		EtherType:   etherType,
		Payload:     []byte(msg),
	}

	b, err := f.MarshalBinary()
	if err != nil {
		log.Fatalf("failed to marshal ethernet frame: %v", err)
	}

	addr := &raw.Addr{
		HardwareAddr: ethernet.Broadcast,
	}

	// Send message forever.
	t := time.NewTicker(1 * time.Second)
	for range t.C {
		if _, err := conn.WriteTo(b, addr); err != nil {
			log.Fatalf("Failed to send message: %v", err)
		}
	}
}

func sendToAll(conn net.PacketConn, source net.HardwareAddr, etherType int, msg string) {
	// Message is broadcast to all machines in same network segment.
	f := &ethernet.Frame{
		Destination: ethernet.Broadcast,
		Source:      source,
		EtherType:   etherType,
		Payload:     []byte(msg),
	}

	b, err := f.MarshalBinary()
	if err != nil {
		log.Fatalf("Failed to marshal ethernet frame: %v", err)
	}

	// Required by Linux, even though the Ethernet frame has a destination.
	// Unused by BSD.
	addr := addr{Broadcast}
	// Send message forever.
	t := time.NewTicker(1 * time.Second)
	for range t.C {
		if _, err := conn.WriteTo(b, addr); err != nil {
			log.Fatalf("failed to send message: %v", err)
		}
	}
}
