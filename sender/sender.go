package sender

import (
	"github.com/kacmak7/go-p2p-packets/protocol"
	"github.com/mdlayher/ethernet"
	"log"
	"net"
)

// sends frame
func send(conn net.PacketConn, dest net.HardwareAddr, source net.HardwareAddr, ether ethernet.EtherType, msg string) {
	f := &ethernet.Frame{
		Destination: dest,
		Source:      source,
		EtherType:   ether,
		Payload:     []byte(msg),
	}

	b, err := f.MarshalBinary()
	if err != nil {
		log.Fatalf("Failed to marshal protocol frame: %v", err)
	}

	addr := &protocol.Addr{
		HardwareAddr: dest,
	}

	if _, err := conn.WriteTo(b, addr); err != nil {
		log.Fatalf("Failed to send message: %v", err)
	}
}

// sends message and waits for the confirmation
func SendMsg(conn net.PacketConn, dest net.HardwareAddr, source net.HardwareAddr, msg string) {
	send(conn, dest, source, protocol.EtherType, msg)
}

// scans network for available listening devices
func Scan(conn net.PacketConn, source net.HardwareAddr) {
	send(conn, ethernet.Broadcast, source, protocol.EtherType,"")
}