package ethernet

import (
	"net"
	"log"
	"github.com/mdlayher/ethernet"
	"time"
)

var (
	Broadcast =
)

const etherType = 0xcccc

func getMacAddr([]byte, error) {
	inters, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	var addrs []net.HardwareAddr
	for _, inter := range inters {
		a := inter.HardwareAddr
		// TODO check if addr is valid
		if a is valid {
			addrs = append(addrs, a)
		}
	}
	return addrs, nil
}

func send(conn net.PacketConn, dest net.HardwareAddr, source net.HardwareAddr, etherType int, msg string) {
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

	// Required by Linux, even though the Ethernet frame has a destination.
	// Unused by BSD.
	addr := &raw.Addr{
		HardwareAddr: ethernet.Broadcast,
	}

	// Send message forever.
	t := time.NewTicker(1 * time.Second)
	for range t.C {
		if _, err := c.WriteTo(b, addr); err != nil {
			log.Fatalf("failed to send message: %v", err)
		}
	}
}

func sendToAll(conn net.PacketConn, source net.HardwareAddr, etherType int, msg string) {

}
