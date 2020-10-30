package ethernet

import (
	"github.com/mdlayher/ethernet"
	"github.com/mdlayher/raw"
	"log"
	"net"
)

const EtherType ethernet.EtherType = 0xcccc

var Broadcast net.HardwareAddr = []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff}

type Addr net.HardwareAddr // TODO whats with this Addr ??

func (Addr) Network() string { return "ethernet" }
func (Addr) String() string { return string(Broadcast) }

// sends frame to the only one node
func Send(conn net.PacketConn, dest net.HardwareAddr, source net.HardwareAddr, msg string) {
	f := &ethernet.Frame{
		Destination: dest,
		Source:      source,
		EtherType:   EtherType,
		Payload:     []byte(msg),
	}

	b, err := f.MarshalBinary()
	if err != nil {
		log.Fatalf("Failed to marshal ethernet frame: %v", err)
	}

	// Required by Linux, even though the Ethernet frame has a destination.
	// Unused by BSD.
	addr := &raw.Addr{
		HardwareAddr: dest,
	}

	if _, err := conn.WriteTo(b, addr); err != nil {
		log.Fatalf("Failed to send message: %v", err)
	}
}
