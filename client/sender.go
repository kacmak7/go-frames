package sender

import (
	protocol "github.com/kacmak7/go-p2p-packets/ethernet"
	"github.com/mdlayher/ethernet"
	"github.com/mdlayher/raw" // TODO
	"log"
	"net"
)

// sends message frame to the another node
func Send(conn net.PacketConn, dest net.HardwareAddr, source net.HardwareAddr, ether ethernet.EtherType, msg string) {
	f := &ethernet.Frame{
		Destination: dest,
		Source:      source,
		EtherType:   ether,
		Payload:     []byte(msg),
	}

	b, err := f.MarshalBinary()
	if err != nil {
		log.Fatalf("Failed to marshal ethernet frame: %v", err)
	}

	addr := &raw.Addr{ // TODO change
		HardwareAddr: dest,
	}

	if _, err := conn.WriteTo(b, addr); err != nil {
		log.Fatalf("Failed to send message: %v", err)
	}
}

// scans network for available listening devices
func Scan(conn net.PacketConn, source net.HardwareAddr) {
	Send(conn, ethernet.Broadcast, source, protocol.EtherType,"scanning")
}