package main

import (
	"github.com/kacmak7/go-p2p-packets/ethernet"
	"github.com/mdlayher/raw"
	"log"
	"net"
)

const (
	INTERFACE = "enp39s0"
	ETHER = 0x22F0
)

var dest net.HardwareAddr = []byte{0x30, 0x9c, 0x23, 0x0e, 0x1a, 0x00}

func main() {
	ifi, err := net.InterfaceByName(INTERFACE)
	if err != nil {
		log.Fatal("interface error")
	}

	c, err := raw.ListenPacket(ifi, ETHER, nil)
	if err != nil {
		log.Fatal("listenpacket error")
	}

	// LOGS
	log.Printf("local %v", ifi.HardwareAddr)
	log.Printf("dest  %v", dest)

	ethernet.Send(c, dest, ifi.HardwareAddr, "message to you")
}
