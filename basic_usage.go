package main

import (
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

func main() {

	myPacketData := []byte{1,2,5}

// Decode a packet
packet := gopacket.NewPacket(myPacketData, layers.LayerTypeEthernet, gopacket.Default)
// Get the TCP layer from this packet
if tcpLayer := packet.Layer(layers.LayerTypeTCP); tcpLayer != nil {
  fmt.Println("This is a TCP packet!")
  // Get actual TCP data from this layer
  tcp, _ := tcpLayer.(*layers.TCP)
  fmt.Printf("From src port %d to dst port %d\n", tcp.SrcPort, tcp.DstPort)
}
// Iterate over all layers, printing out each layer type
for _, layer := range packet.Layers() {
  fmt.Println("PACKET LAYER:", layer.LayerType())
}
}
