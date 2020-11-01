package sender

import (
	protocol "github.com/kacmak7/go-p2p-packets/ethernet"
	"github.com/mdlayher/ethernet"
	"github.com/mdlayher/raw" // TODO
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
		log.Fatalf("Failed to marshal ethernet frame: %v", err)
	}

	addr := &raw.Addr{ // TODO change
		HardwareAddr: dest,
	}

	if _, err := conn.WriteTo(b, addr); err != nil {
		log.Fatalf("Failed to send message: %v", err)
	}
}

// Sends ping and waits for pong
func ping(conn net.PacketConn, dest net.HardwareAddr, source net.HardwareAddr, msg string) {
	send(conn, dest, source, protocol.EtherType, protocol.PingMsg)
}

func SendMsg(conn net.PacketConn, dest net.HardwareAddr, source net.HardwareAddr, msg string) {
	send(conn, dest, source, protocol.EtherType, msg)
}

// scans network for available listening devices
// scanning
func Scan(conn net.PacketConn, source net.HardwareAddr) {
	send(conn, ethernet.Broadcast, source, protocol.EtherType,"")
}